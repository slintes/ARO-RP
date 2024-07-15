// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/frontend/adminactions (interfaces: KubeActions,AzureActions,AppLensActions)

// Package mock_adminactions is a generated GoMock package.
package mock_adminactions

import (
	context "context"
	io "io"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	features "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
)

// MockKubeActions is a mock of KubeActions interface.
type MockKubeActions struct {
	ctrl     *gomock.Controller
	recorder *MockKubeActionsMockRecorder
}

// MockKubeActionsMockRecorder is the mock recorder for MockKubeActions.
type MockKubeActionsMockRecorder struct {
	mock *MockKubeActions
}

// NewMockKubeActions creates a new mock instance.
func NewMockKubeActions(ctrl *gomock.Controller) *MockKubeActions {
	mock := &MockKubeActions{ctrl: ctrl}
	mock.recorder = &MockKubeActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeActions) EXPECT() *MockKubeActionsMockRecorder {
	return m.recorder
}

// ApproveAllCsrs mocks base method.
func (m *MockKubeActions) ApproveAllCsrs(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveAllCsrs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveAllCsrs indicates an expected call of ApproveAllCsrs.
func (mr *MockKubeActionsMockRecorder) ApproveAllCsrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveAllCsrs", reflect.TypeOf((*MockKubeActions)(nil).ApproveAllCsrs), arg0)
}

// ApproveCsr mocks base method.
func (m *MockKubeActions) ApproveCsr(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveCsr", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveCsr indicates an expected call of ApproveCsr.
func (mr *MockKubeActionsMockRecorder) ApproveCsr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveCsr", reflect.TypeOf((*MockKubeActions)(nil).ApproveCsr), arg0, arg1)
}

// CordonNode mocks base method.
func (m *MockKubeActions) CordonNode(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CordonNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CordonNode indicates an expected call of CordonNode.
func (mr *MockKubeActionsMockRecorder) CordonNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CordonNode", reflect.TypeOf((*MockKubeActions)(nil).CordonNode), arg0, arg1, arg2)
}

// DrainNode mocks base method.
func (m *MockKubeActions) DrainNode(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrainNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DrainNode indicates an expected call of DrainNode.
func (mr *MockKubeActionsMockRecorder) DrainNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrainNode", reflect.TypeOf((*MockKubeActions)(nil).DrainNode), arg0, arg1)
}

// KubeCreateOrUpdate mocks base method.
func (m *MockKubeActions) KubeCreateOrUpdate(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeCreateOrUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// KubeCreateOrUpdate indicates an expected call of KubeCreateOrUpdate.
func (mr *MockKubeActionsMockRecorder) KubeCreateOrUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeCreateOrUpdate", reflect.TypeOf((*MockKubeActions)(nil).KubeCreateOrUpdate), arg0, arg1)
}

// KubeDelete mocks base method.
func (m *MockKubeActions) KubeDelete(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool, arg5 *v1.DeletionPropagation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeDelete", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// KubeDelete indicates an expected call of KubeDelete.
func (mr *MockKubeActionsMockRecorder) KubeDelete(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeDelete", reflect.TypeOf((*MockKubeActions)(nil).KubeDelete), arg0, arg1, arg2, arg3, arg4, arg5)
}

// KubeGet mocks base method.
func (m *MockKubeActions) KubeGet(arg0 context.Context, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeGet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeGet indicates an expected call of KubeGet.
func (mr *MockKubeActionsMockRecorder) KubeGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeGet", reflect.TypeOf((*MockKubeActions)(nil).KubeGet), arg0, arg1, arg2, arg3)
}

// KubeGetPodLogs mocks base method.
func (m *MockKubeActions) KubeGetPodLogs(arg0 context.Context, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeGetPodLogs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeGetPodLogs indicates an expected call of KubeGetPodLogs.
func (mr *MockKubeActionsMockRecorder) KubeGetPodLogs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeGetPodLogs", reflect.TypeOf((*MockKubeActions)(nil).KubeGetPodLogs), arg0, arg1, arg2, arg3)
}

// KubeList mocks base method.
func (m *MockKubeActions) KubeList(arg0 context.Context, arg1, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeList indicates an expected call of KubeList.
func (mr *MockKubeActionsMockRecorder) KubeList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeList", reflect.TypeOf((*MockKubeActions)(nil).KubeList), arg0, arg1, arg2)
}

// KubeWatch mocks base method.
func (m *MockKubeActions) KubeWatch(arg0 context.Context, arg1 *unstructured.Unstructured, arg2 string) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeWatch", arg0, arg1, arg2)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeWatch indicates an expected call of KubeWatch.
func (mr *MockKubeActionsMockRecorder) KubeWatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeWatch", reflect.TypeOf((*MockKubeActions)(nil).KubeWatch), arg0, arg1, arg2)
}

// ResolveGVR mocks base method.
func (m *MockKubeActions) ResolveGVR(arg0, arg1 string) (schema.GroupVersionResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveGVR", arg0, arg1)
	ret0, _ := ret[0].(schema.GroupVersionResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveGVR indicates an expected call of ResolveGVR.
func (mr *MockKubeActionsMockRecorder) ResolveGVR(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveGVR", reflect.TypeOf((*MockKubeActions)(nil).ResolveGVR), arg0, arg1)
}

// MockAzureActions is a mock of AzureActions interface.
type MockAzureActions struct {
	ctrl     *gomock.Controller
	recorder *MockAzureActionsMockRecorder
}

// MockAzureActionsMockRecorder is the mock recorder for MockAzureActions.
type MockAzureActionsMockRecorder struct {
	mock *MockAzureActions
}

// NewMockAzureActions creates a new mock instance.
func NewMockAzureActions(ctrl *gomock.Controller) *MockAzureActions {
	mock := &MockAzureActions{ctrl: ctrl}
	mock.recorder = &MockAzureActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzureActions) EXPECT() *MockAzureActionsMockRecorder {
	return m.recorder
}

// GroupResourceList mocks base method.
func (m *MockAzureActions) GroupResourceList(arg0 context.Context) ([]features.GenericResourceExpanded, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupResourceList", arg0)
	ret0, _ := ret[0].([]features.GenericResourceExpanded)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupResourceList indicates an expected call of GroupResourceList.
func (mr *MockAzureActionsMockRecorder) GroupResourceList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupResourceList", reflect.TypeOf((*MockAzureActions)(nil).GroupResourceList), arg0)
}

// NICReconcileFailedState mocks base method.
func (m *MockAzureActions) NICReconcileFailedState(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NICReconcileFailedState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NICReconcileFailedState indicates an expected call of NICReconcileFailedState.
func (mr *MockAzureActionsMockRecorder) NICReconcileFailedState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NICReconcileFailedState", reflect.TypeOf((*MockAzureActions)(nil).NICReconcileFailedState), arg0, arg1)
}

// ResourceDeleteAndWait mocks base method.
func (m *MockAzureActions) ResourceDeleteAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceDeleteAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceDeleteAndWait indicates an expected call of ResourceDeleteAndWait.
func (mr *MockAzureActionsMockRecorder) ResourceDeleteAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceDeleteAndWait", reflect.TypeOf((*MockAzureActions)(nil).ResourceDeleteAndWait), arg0, arg1)
}

// ResourceGroupHasVM mocks base method.
func (m *MockAzureActions) ResourceGroupHasVM(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroupHasVM", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceGroupHasVM indicates an expected call of ResourceGroupHasVM.
func (mr *MockAzureActionsMockRecorder) ResourceGroupHasVM(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroupHasVM", reflect.TypeOf((*MockAzureActions)(nil).ResourceGroupHasVM), arg0, arg1)
}

// ResourcesList mocks base method.
func (m *MockAzureActions) ResourcesList(arg0 context.Context, arg1 []features.GenericResourceExpanded, arg2 io.WriteCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourcesList", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourcesList indicates an expected call of ResourcesList.
func (mr *MockAzureActionsMockRecorder) ResourcesList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourcesList", reflect.TypeOf((*MockAzureActions)(nil).ResourcesList), arg0, arg1, arg2)
}

// VMRedeployAndWait mocks base method.
func (m *MockAzureActions) VMRedeployAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMRedeployAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMRedeployAndWait indicates an expected call of VMRedeployAndWait.
func (mr *MockAzureActionsMockRecorder) VMRedeployAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMRedeployAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMRedeployAndWait), arg0, arg1)
}

// VMResize mocks base method.
func (m *MockAzureActions) VMResize(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMResize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMResize indicates an expected call of VMResize.
func (mr *MockAzureActionsMockRecorder) VMResize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMResize", reflect.TypeOf((*MockAzureActions)(nil).VMResize), arg0, arg1, arg2)
}

// VMSerialConsole mocks base method.
func (m *MockAzureActions) VMSerialConsole(arg0 context.Context, arg1 *logrus.Entry, arg2 string, arg3 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSerialConsole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMSerialConsole indicates an expected call of VMSerialConsole.
func (mr *MockAzureActionsMockRecorder) VMSerialConsole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSerialConsole", reflect.TypeOf((*MockAzureActions)(nil).VMSerialConsole), arg0, arg1, arg2, arg3)
}

// VMSizeList mocks base method.
func (m *MockAzureActions) VMSizeList(arg0 context.Context) ([]compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSizeList", arg0)
	ret0, _ := ret[0].([]compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VMSizeList indicates an expected call of VMSizeList.
func (mr *MockAzureActionsMockRecorder) VMSizeList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSizeList", reflect.TypeOf((*MockAzureActions)(nil).VMSizeList), arg0)
}

// VMStartAndWait mocks base method.
func (m *MockAzureActions) VMStartAndWait(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMStartAndWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMStartAndWait indicates an expected call of VMStartAndWait.
func (mr *MockAzureActionsMockRecorder) VMStartAndWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMStartAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMStartAndWait), arg0, arg1)
}

// VMStopAndWait mocks base method.
func (m *MockAzureActions) VMStopAndWait(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMStopAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VMStopAndWait indicates an expected call of VMStopAndWait.
func (mr *MockAzureActionsMockRecorder) VMStopAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMStopAndWait", reflect.TypeOf((*MockAzureActions)(nil).VMStopAndWait), arg0, arg1, arg2)
}

// WriteToStream mocks base method.
func (m *MockAzureActions) WriteToStream(arg0 context.Context, arg1 io.WriteCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToStream indicates an expected call of WriteToStream.
func (mr *MockAzureActionsMockRecorder) WriteToStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToStream", reflect.TypeOf((*MockAzureActions)(nil).WriteToStream), arg0, arg1)
}

// MockAppLensActions is a mock of AppLensActions interface.
type MockAppLensActions struct {
	ctrl     *gomock.Controller
	recorder *MockAppLensActionsMockRecorder
}

// MockAppLensActionsMockRecorder is the mock recorder for MockAppLensActions.
type MockAppLensActionsMockRecorder struct {
	mock *MockAppLensActions
}

// NewMockAppLensActions creates a new mock instance.
func NewMockAppLensActions(ctrl *gomock.Controller) *MockAppLensActions {
	mock := &MockAppLensActions{ctrl: ctrl}
	mock.recorder = &MockAppLensActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLensActions) EXPECT() *MockAppLensActionsMockRecorder {
	return m.recorder
}

// AppLensGetDetector mocks base method.
func (m *MockAppLensActions) AppLensGetDetector(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppLensGetDetector", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppLensGetDetector indicates an expected call of AppLensGetDetector.
func (mr *MockAppLensActionsMockRecorder) AppLensGetDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppLensGetDetector", reflect.TypeOf((*MockAppLensActions)(nil).AppLensGetDetector), arg0, arg1)
}

// AppLensListDetectors mocks base method.
func (m *MockAppLensActions) AppLensListDetectors(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppLensListDetectors", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppLensListDetectors indicates an expected call of AppLensListDetectors.
func (mr *MockAppLensActionsMockRecorder) AppLensListDetectors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppLensListDetectors", reflect.TypeOf((*MockAppLensActions)(nil).AppLensListDetectors), arg0)
}
