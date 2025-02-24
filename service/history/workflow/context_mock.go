// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package workflow is a generated GoMock package.
package workflow

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/server/api/persistence/v1"
	namespace "go.temporal.io/server/common/namespace"
	persistence "go.temporal.io/server/common/persistence"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockContext) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockContextMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockContext)(nil).Clear))
}

// ConflictResolveWorkflowExecution mocks base method.
func (m *MockContext) ConflictResolveWorkflowExecution(ctx context.Context, now time.Time, conflictResolveMode persistence.ConflictResolveWorkflowMode, resetMutableState MutableState, newContext Context, newMutableState MutableState, currentContext Context, currentMutableState MutableState, currentTransactionPolicy *TransactionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConflictResolveWorkflowExecution", ctx, now, conflictResolveMode, resetMutableState, newContext, newMutableState, currentContext, currentMutableState, currentTransactionPolicy)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConflictResolveWorkflowExecution indicates an expected call of ConflictResolveWorkflowExecution.
func (mr *MockContextMockRecorder) ConflictResolveWorkflowExecution(ctx, now, conflictResolveMode, resetMutableState, newContext, newMutableState, currentContext, currentMutableState, currentTransactionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConflictResolveWorkflowExecution", reflect.TypeOf((*MockContext)(nil).ConflictResolveWorkflowExecution), ctx, now, conflictResolveMode, resetMutableState, newContext, newMutableState, currentContext, currentMutableState, currentTransactionPolicy)
}

// CreateWorkflowExecution mocks base method.
func (m *MockContext) CreateWorkflowExecution(ctx context.Context, now time.Time, createMode persistence.CreateWorkflowMode, prevRunID string, prevLastWriteVersion int64, newMutableState MutableState, newWorkflow *persistence.WorkflowSnapshot, newWorkflowEvents []*persistence.WorkflowEvents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkflowExecution", ctx, now, createMode, prevRunID, prevLastWriteVersion, newMutableState, newWorkflow, newWorkflowEvents)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorkflowExecution indicates an expected call of CreateWorkflowExecution.
func (mr *MockContextMockRecorder) CreateWorkflowExecution(ctx, now, createMode, prevRunID, prevLastWriteVersion, newMutableState, newWorkflow, newWorkflowEvents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflowExecution", reflect.TypeOf((*MockContext)(nil).CreateWorkflowExecution), ctx, now, createMode, prevRunID, prevLastWriteVersion, newMutableState, newWorkflow, newWorkflowEvents)
}

// GetHistorySize mocks base method.
func (m *MockContext) GetHistorySize() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistorySize")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetHistorySize indicates an expected call of GetHistorySize.
func (mr *MockContextMockRecorder) GetHistorySize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistorySize", reflect.TypeOf((*MockContext)(nil).GetHistorySize))
}

// GetNamespace mocks base method.
func (m *MockContext) GetNamespace() namespace.Name {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(namespace.Name)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockContextMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockContext)(nil).GetNamespace))
}

// GetNamespaceID mocks base method.
func (m *MockContext) GetNamespaceID() namespace.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceID")
	ret0, _ := ret[0].(namespace.ID)
	return ret0
}

// GetNamespaceID indicates an expected call of GetNamespaceID.
func (mr *MockContextMockRecorder) GetNamespaceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceID", reflect.TypeOf((*MockContext)(nil).GetNamespaceID))
}

// GetRunID mocks base method.
func (m *MockContext) GetRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunID indicates an expected call of GetRunID.
func (mr *MockContextMockRecorder) GetRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunID", reflect.TypeOf((*MockContext)(nil).GetRunID))
}

// GetWorkflowID mocks base method.
func (m *MockContext) GetWorkflowID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkflowID indicates an expected call of GetWorkflowID.
func (mr *MockContextMockRecorder) GetWorkflowID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowID", reflect.TypeOf((*MockContext)(nil).GetWorkflowID))
}

// LoadExecutionStats mocks base method.
func (m *MockContext) LoadExecutionStats(ctx context.Context) (*v1.ExecutionStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadExecutionStats", ctx)
	ret0, _ := ret[0].(*v1.ExecutionStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadExecutionStats indicates an expected call of LoadExecutionStats.
func (mr *MockContextMockRecorder) LoadExecutionStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadExecutionStats", reflect.TypeOf((*MockContext)(nil).LoadExecutionStats), ctx)
}

// LoadMutableState mocks base method.
func (m *MockContext) LoadMutableState(ctx context.Context) (MutableState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMutableState", ctx)
	ret0, _ := ret[0].(MutableState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMutableState indicates an expected call of LoadMutableState.
func (mr *MockContextMockRecorder) LoadMutableState(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMutableState", reflect.TypeOf((*MockContext)(nil).LoadMutableState), ctx)
}

// Lock mocks base method.
func (m *MockContext) Lock(ctx context.Context, caller CallerType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, caller)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock.
func (mr *MockContextMockRecorder) Lock(ctx, caller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockContext)(nil).Lock), ctx, caller)
}

// PersistWorkflowEvents mocks base method.
func (m *MockContext) PersistWorkflowEvents(ctx context.Context, workflowEvents *persistence.WorkflowEvents) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistWorkflowEvents", ctx, workflowEvents)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PersistWorkflowEvents indicates an expected call of PersistWorkflowEvents.
func (mr *MockContextMockRecorder) PersistWorkflowEvents(ctx, workflowEvents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistWorkflowEvents", reflect.TypeOf((*MockContext)(nil).PersistWorkflowEvents), ctx, workflowEvents)
}

// ReapplyEvents mocks base method.
func (m *MockContext) ReapplyEvents(eventBatches []*persistence.WorkflowEvents) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReapplyEvents", eventBatches)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReapplyEvents indicates an expected call of ReapplyEvents.
func (mr *MockContextMockRecorder) ReapplyEvents(eventBatches interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapplyEvents", reflect.TypeOf((*MockContext)(nil).ReapplyEvents), eventBatches)
}

// SetHistorySize mocks base method.
func (m *MockContext) SetHistorySize(size int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHistorySize", size)
}

// SetHistorySize indicates an expected call of SetHistorySize.
func (mr *MockContextMockRecorder) SetHistorySize(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHistorySize", reflect.TypeOf((*MockContext)(nil).SetHistorySize), size)
}

// SetWorkflowExecution mocks base method.
func (m *MockContext) SetWorkflowExecution(ctx context.Context, now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkflowExecution", ctx, now)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkflowExecution indicates an expected call of SetWorkflowExecution.
func (mr *MockContextMockRecorder) SetWorkflowExecution(ctx, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkflowExecution", reflect.TypeOf((*MockContext)(nil).SetWorkflowExecution), ctx, now)
}

// Unlock mocks base method.
func (m *MockContext) Unlock(caller CallerType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock", caller)
}

// Unlock indicates an expected call of Unlock.
func (mr *MockContextMockRecorder) Unlock(caller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockContext)(nil).Unlock), caller)
}

// UpdateWorkflowExecutionAsActive mocks base method.
func (m *MockContext) UpdateWorkflowExecutionAsActive(ctx context.Context, now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecutionAsActive", ctx, now)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowExecutionAsActive indicates an expected call of UpdateWorkflowExecutionAsActive.
func (mr *MockContextMockRecorder) UpdateWorkflowExecutionAsActive(ctx, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecutionAsActive", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecutionAsActive), ctx, now)
}

// UpdateWorkflowExecutionAsPassive mocks base method.
func (m *MockContext) UpdateWorkflowExecutionAsPassive(ctx context.Context, now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecutionAsPassive", ctx, now)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowExecutionAsPassive indicates an expected call of UpdateWorkflowExecutionAsPassive.
func (mr *MockContextMockRecorder) UpdateWorkflowExecutionAsPassive(ctx, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecutionAsPassive", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecutionAsPassive), ctx, now)
}

// UpdateWorkflowExecutionWithNew mocks base method.
func (m *MockContext) UpdateWorkflowExecutionWithNew(ctx context.Context, now time.Time, updateMode persistence.UpdateWorkflowMode, newContext Context, newMutableState MutableState, currentWorkflowTransactionPolicy TransactionPolicy, newWorkflowTransactionPolicy *TransactionPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecutionWithNew", ctx, now, updateMode, newContext, newMutableState, currentWorkflowTransactionPolicy, newWorkflowTransactionPolicy)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowExecutionWithNew indicates an expected call of UpdateWorkflowExecutionWithNew.
func (mr *MockContextMockRecorder) UpdateWorkflowExecutionWithNew(ctx, now, updateMode, newContext, newMutableState, currentWorkflowTransactionPolicy, newWorkflowTransactionPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecutionWithNew", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecutionWithNew), ctx, now, updateMode, newContext, newMutableState, currentWorkflowTransactionPolicy, newWorkflowTransactionPolicy)
}

// UpdateWorkflowExecutionWithNewAsActive mocks base method.
func (m *MockContext) UpdateWorkflowExecutionWithNewAsActive(ctx context.Context, now time.Time, newContext Context, newMutableState MutableState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecutionWithNewAsActive", ctx, now, newContext, newMutableState)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowExecutionWithNewAsActive indicates an expected call of UpdateWorkflowExecutionWithNewAsActive.
func (mr *MockContextMockRecorder) UpdateWorkflowExecutionWithNewAsActive(ctx, now, newContext, newMutableState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecutionWithNewAsActive", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecutionWithNewAsActive), ctx, now, newContext, newMutableState)
}

// UpdateWorkflowExecutionWithNewAsPassive mocks base method.
func (m *MockContext) UpdateWorkflowExecutionWithNewAsPassive(ctx context.Context, now time.Time, newContext Context, newMutableState MutableState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecutionWithNewAsPassive", ctx, now, newContext, newMutableState)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowExecutionWithNewAsPassive indicates an expected call of UpdateWorkflowExecutionWithNewAsPassive.
func (mr *MockContextMockRecorder) UpdateWorkflowExecutionWithNewAsPassive(ctx, now, newContext, newMutableState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecutionWithNewAsPassive", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecutionWithNewAsPassive), ctx, now, newContext, newMutableState)
}
