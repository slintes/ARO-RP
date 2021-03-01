package validate

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/api/validate/dynamic"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/aad"
	"github.com/Azure/ARO-RP/pkg/util/refreshable"
	"github.com/Azure/ARO-RP/pkg/util/subnet"
)

// OpenShiftClusterDynamicValidator is the dynamic validator interface
type OpenShiftClusterDynamicValidator interface {
	Dynamic(context.Context) error
}

// NewOpenShiftClusterDynamicValidator creates a new OpenShiftClusterDynamicValidator
func NewOpenShiftClusterDynamicValidator(log *logrus.Entry, env env.Core, oc *api.OpenShiftCluster, subscriptionDoc *api.SubscriptionDocument, fpAuthorizer refreshable.Authorizer) OpenShiftClusterDynamicValidator {
	return &openShiftClusterDynamicValidator{
		log: log,
		env: env,

		oc:              oc,
		subscriptionDoc: subscriptionDoc,
		fpAuthorizer:    fpAuthorizer,
	}
}

type openShiftClusterDynamicValidator struct {
	log *logrus.Entry
	env env.Core

	oc              *api.OpenShiftCluster
	subscriptionDoc *api.SubscriptionDocument
	fpAuthorizer    refreshable.Authorizer
}

// Dynamic validates an OpenShift cluster
func (dv *openShiftClusterDynamicValidator) Dynamic(ctx context.Context) error {
	// Get all subnets
	subnetIDs := []string{dv.oc.Properties.MasterProfile.SubnetID}

	for _, s := range dv.oc.Properties.WorkerProfiles {
		subnetIDs = append(subnetIDs, s.SubnetID)
	}

	vnetID, _, err := subnet.Split(subnetIDs[0])
	if err != nil {
		return err
	}

	// FP validation
	fpDynamic, err := dynamic.NewValidator(dv.log, dv.env.Environment(), dv.subscriptionDoc.ID, dv.fpAuthorizer, dynamic.AuthorizerFirstParty)
	if err != nil {
		return err
	}

	err = fpDynamic.ValidateVnetPermissions(ctx, vnetID)
	if err != nil {
		return err
	}

	err = fpDynamic.ValidateRouteTablesPermissions(ctx, subnetIDs)
	if err != nil {
		return err
	}

	spp := dv.oc.Properties.ServicePrincipalProfile
	token, err := aad.GetToken(ctx, dv.log, spp.ClientID, string(spp.ClientSecret), dv.subscriptionDoc.Subscription.Properties.TenantID, dv.env.Environment().ActiveDirectoryEndpoint, dv.env.Environment().ResourceManagerEndpoint)
	if err != nil {
		return err
	}

	spAuthorizer := refreshable.NewAuthorizer(token)

	spDynamic, err := dynamic.NewValidator(dv.log, dv.env.Environment(), dv.subscriptionDoc.ID, spAuthorizer, dynamic.AuthorizerClusterServicePrincipal)
	if err != nil {
		return err
	}

	// SP validation
	err = spDynamic.ValidateClusterServicePrincipalProfile(ctx, spp.ClientID, string(spp.ClientSecret), dv.subscriptionDoc.Subscription.Properties.TenantID)
	if err != nil {
		return err
	}

	err = spDynamic.ValidateVnetPermissions(ctx, vnetID)
	if err != nil {
		return err
	}

	err = spDynamic.ValidateRouteTablesPermissions(ctx, subnetIDs)
	if err != nil {
		return err
	}

	// Additional checks - use any dynamic because they both have the correct permissions
	err = spDynamic.ValidateCIDRRanges(ctx, subnetIDs, dv.oc.Properties.NetworkProfile.PodCIDR, dv.oc.Properties.NetworkProfile.ServiceCIDR)
	if err != nil {
		return err
	}

	err = spDynamic.ValidateVnetLocation(ctx, vnetID, dv.oc.Location)
	if err != nil {
		return err
	}

	err = spDynamic.ValidateProviders(ctx)
	if err != nil {
		return err
	}

	err = spDynamic.ValidateQuota(ctx, dv.oc)
	if err != nil {
		return err
	}

	return nil
}
