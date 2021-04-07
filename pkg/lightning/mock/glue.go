// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/br/pkg/lightning/glue/glue.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	parser "github.com/pingcap/parser"
	model "github.com/pingcap/parser/model"
	checkpoints "github.com/pingcap/br/pkg/lightning/checkpoints"
	config "github.com/pingcap/br/pkg/lightning/config"
	glue "github.com/pingcap/br/pkg/lightning/glue"
	log "github.com/pingcap/br/pkg/lightning/log"
	reflect "reflect"
)

// MockGlue is a mock of Glue interface
type MockGlue struct {
	ctrl     *gomock.Controller
	recorder *MockGlueMockRecorder
}

// MockGlueMockRecorder is the mock recorder for MockGlue
type MockGlueMockRecorder struct {
	mock *MockGlue
}

// NewMockGlue creates a new mock instance
func NewMockGlue(ctrl *gomock.Controller) *MockGlue {
	mock := &MockGlue{ctrl: ctrl}
	mock.recorder = &MockGlueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGlue) EXPECT() *MockGlueMockRecorder {
	return m.recorder
}

// OwnsSQLExecutor mocks base method
func (m *MockGlue) OwnsSQLExecutor() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OwnsSQLExecutor")
	ret0, _ := ret[0].(bool)
	return ret0
}

// OwnsSQLExecutor indicates an expected call of OwnsSQLExecutor
func (mr *MockGlueMockRecorder) OwnsSQLExecutor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OwnsSQLExecutor", reflect.TypeOf((*MockGlue)(nil).OwnsSQLExecutor))
}

// GetSQLExecutor mocks base method
func (m *MockGlue) GetSQLExecutor() glue.SQLExecutor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSQLExecutor")
	ret0, _ := ret[0].(glue.SQLExecutor)
	return ret0
}

// GetSQLExecutor indicates an expected call of GetSQLExecutor
func (mr *MockGlueMockRecorder) GetSQLExecutor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSQLExecutor", reflect.TypeOf((*MockGlue)(nil).GetSQLExecutor))
}

// GetDB mocks base method
func (m *MockGlue) GetDB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDB indicates an expected call of GetDB
func (mr *MockGlueMockRecorder) GetDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockGlue)(nil).GetDB))
}

// GetParser mocks base method
func (m *MockGlue) GetParser() *parser.Parser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParser")
	ret0, _ := ret[0].(*parser.Parser)
	return ret0
}

// GetParser indicates an expected call of GetParser
func (mr *MockGlueMockRecorder) GetParser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParser", reflect.TypeOf((*MockGlue)(nil).GetParser))
}

// GetTables mocks base method
func (m *MockGlue) GetTables(arg0 context.Context, arg1 string) ([]*model.TableInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTables", arg0, arg1)
	ret0, _ := ret[0].([]*model.TableInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTables indicates an expected call of GetTables
func (mr *MockGlueMockRecorder) GetTables(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockGlue)(nil).GetTables), arg0, arg1)
}

// GetSession mocks base method
func (m *MockGlue) GetSession(arg0 context.Context) (checkpoints.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0)
	ret0, _ := ret[0].(checkpoints.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockGlueMockRecorder) GetSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockGlue)(nil).GetSession), arg0)
}

// OpenCheckpointsDB mocks base method
func (m *MockGlue) OpenCheckpointsDB(arg0 context.Context, arg1 *config.Config) (checkpoints.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenCheckpointsDB", arg0, arg1)
	ret0, _ := ret[0].(checkpoints.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenCheckpointsDB indicates an expected call of OpenCheckpointsDB
func (mr *MockGlueMockRecorder) OpenCheckpointsDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenCheckpointsDB", reflect.TypeOf((*MockGlue)(nil).OpenCheckpointsDB), arg0, arg1)
}

// Record mocks base method
func (m *MockGlue) Record(arg0 string, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Record", arg0, arg1)
}

// Record indicates an expected call of Record
func (mr *MockGlueMockRecorder) Record(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockGlue)(nil).Record), arg0, arg1)
}

// MockSQLExecutor is a mock of SQLExecutor interface
type MockSQLExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockSQLExecutorMockRecorder
}

// MockSQLExecutorMockRecorder is the mock recorder for MockSQLExecutor
type MockSQLExecutorMockRecorder struct {
	mock *MockSQLExecutor
}

// NewMockSQLExecutor creates a new mock instance
func NewMockSQLExecutor(ctrl *gomock.Controller) *MockSQLExecutor {
	mock := &MockSQLExecutor{ctrl: ctrl}
	mock.recorder = &MockSQLExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSQLExecutor) EXPECT() *MockSQLExecutorMockRecorder {
	return m.recorder
}

// ExecuteWithLog mocks base method
func (m *MockSQLExecutor) ExecuteWithLog(ctx context.Context, query, purpose string, logger log.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteWithLog", ctx, query, purpose, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteWithLog indicates an expected call of ExecuteWithLog
func (mr *MockSQLExecutorMockRecorder) ExecuteWithLog(ctx, query, purpose, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteWithLog", reflect.TypeOf((*MockSQLExecutor)(nil).ExecuteWithLog), ctx, query, purpose, logger)
}

// ObtainStringWithLog mocks base method
func (m *MockSQLExecutor) ObtainStringWithLog(ctx context.Context, query, purpose string, logger log.Logger) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObtainStringWithLog", ctx, query, purpose, logger)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtainStringWithLog indicates an expected call of ObtainStringWithLog
func (mr *MockSQLExecutorMockRecorder) ObtainStringWithLog(ctx, query, purpose, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtainStringWithLog", reflect.TypeOf((*MockSQLExecutor)(nil).ObtainStringWithLog), ctx, query, purpose, logger)
}

// QueryStringsWithLog mocks base method
func (m *MockSQLExecutor) QueryStringsWithLog(ctx context.Context, query, purpose string, logger log.Logger) ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStringsWithLog", ctx, query, purpose, logger)
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryStringsWithLog indicates an expected call of QueryStringsWithLog
func (mr *MockSQLExecutorMockRecorder) QueryStringsWithLog(ctx, query, purpose, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStringsWithLog", reflect.TypeOf((*MockSQLExecutor)(nil).QueryStringsWithLog), ctx, query, purpose, logger)
}

// Close mocks base method
func (m *MockSQLExecutor) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSQLExecutorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSQLExecutor)(nil).Close))
}
