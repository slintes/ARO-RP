package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	configv1 "github.com/openshift/api/config/v1"
	configfake "github.com/openshift/client-go/config/clientset/versioned/fake"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/Azure/ARO-RP/pkg/api"
	pkgoperator "github.com/Azure/ARO-RP/pkg/operator"
	mock_metrics "github.com/Azure/ARO-RP/pkg/util/mocks/metrics"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

func TestEmitClusterVersion(t *testing.T) {
	ctx := context.Background()

	cli := fake.NewSimpleClientset(
		&appsv1.Deployment{ // metrics expected
			ObjectMeta: metav1.ObjectMeta{
				Namespace: pkgoperator.Namespace,
				Name:      "aro-operator-master",
				Labels: map[string]string{
					"version": "test",
				},
			},
		},
	)

	for _, tt := range []struct {
		name                                     string
		cv                                       *configv1.ClusterVersion
		oc                                       *api.OpenShiftCluster
		wantActualVersion                        string
		wantDesiredVersion                       string
		wantProvisionedByResourceProviderVersion string
		wantAvailableVersion                     string
	}{
		{
			name: "without spec",
			cv: &configv1.ClusterVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name: "version",
				},
				Status: configv1.ClusterVersionStatus{
					Desired: configv1.Update{
						Version: "4.3.3",
					},
					History: []configv1.UpdateHistory{
						{
							State:   configv1.PartialUpdate,
							Version: "4.3.2",
						},
						{
							State:   configv1.CompletedUpdate,
							Version: "4.3.1",
						},
						{
							State:   configv1.CompletedUpdate,
							Version: "4.3.0",
						},
					},
				},
			},
			oc: &api.OpenShiftCluster{
				Properties: api.OpenShiftClusterProperties{},
			},
			wantActualVersion:                        "4.3.1",
			wantDesiredVersion:                       "4.3.3",
			wantProvisionedByResourceProviderVersion: "",
			wantAvailableVersion:                     "4.3.40",
		},
		{
			name: "with spec",
			cv: &configv1.ClusterVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name: "version",
				},
				Spec: configv1.ClusterVersionSpec{
					DesiredUpdate: &configv1.Update{
						Version: "4.3.4",
					},
				},
				Status: configv1.ClusterVersionStatus{
					Desired: configv1.Update{
						Version: "4.3.3",
					},
				},
			},
			oc: &api.OpenShiftCluster{
				Properties: api.OpenShiftClusterProperties{},
			},
			wantDesiredVersion:                       "4.3.4",
			wantProvisionedByResourceProviderVersion: "",
		},
		{
			name: "with ProvisionedBy",
			cv: &configv1.ClusterVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name: "version",
				},
			},
			oc: &api.OpenShiftCluster{
				Properties: api.OpenShiftClusterProperties{
					ProvisionedBy: "somesha",
				},
			},
			wantProvisionedByResourceProviderVersion: "somesha",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			configcli := configfake.NewSimpleClientset(tt.cv)

			controller := gomock.NewController(t)
			defer controller.Finish()

			m := mock_metrics.NewMockInterface(controller)

			mon := &Monitor{
				configcli: configcli,
				m:         m,
				oc:        tt.oc,
				cli:       cli,
			}

			m.EXPECT().EmitGauge("cluster.versions", int64(1), map[string]string{
				"actualVersion":                        tt.wantActualVersion,
				"desiredVersion":                       tt.wantDesiredVersion,
				"provisionedByResourceProviderVersion": tt.wantProvisionedByResourceProviderVersion,
				"operatorVersion":                      "test",
				"resourceProviderVersion":              "unknown",
				"availableVersion":                     tt.wantAvailableVersion,
			})

			err := mon.emitClusterVersions(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestBaselineVersion(t *testing.T) {
	streams := []*version.Stream{
		{
			Version: version.NewVersion(4, 3, 18),
		},
		{
			Version: version.NewVersion(4, 4, 38),
		},
		{
			Version: version.NewVersion(4, 5, 2),
		},
	}

	for _, tt := range []struct {
		cv      configv1.ClusterVersion
		want    string
		upgrade bool
	}{
		{
			cv: configv1.ClusterVersion{
				Status: configv1.ClusterVersionStatus{
					History: []configv1.UpdateHistory{
						{
							State:   configv1.CompletedUpdate,
							Version: "4.3.2",
						},
					},
				},
			},
			want: "4.3.18",
		},
		{
			cv: configv1.ClusterVersion{
				Status: configv1.ClusterVersionStatus{
					History: []configv1.UpdateHistory{
						{
							State:   configv1.CompletedUpdate,
							Version: "4.4.40",
						},
					},
				},
			},
			want: "",
		},
		{
			cv: configv1.ClusterVersion{
				Status: configv1.ClusterVersionStatus{
					History: []configv1.UpdateHistory{
						{
							State:   configv1.CompletedUpdate,
							Version: "4.5.1",
						},
					},
				},
			},
			want: "4.5.2",
		},
		{
			cv: configv1.ClusterVersion{
				Status: configv1.ClusterVersionStatus{
					History: []configv1.UpdateHistory{
						{
							State:   configv1.CompletedUpdate,
							Version: "4.5.2",
						},
					},
				},
			},
			want: "",
		},
	} {
		bVersion := availableVersion(&tt.cv, streams)
		if bVersion != tt.want {
			t.Fatalf("Upgrade version does not match: %s\n", bVersion)
		}
	}
}
