// Code generated by MockGen. DO NOT EDIT.
// Source: app/testsupport/mock/src.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDataSet is a mock of DataSet interface
type MockDataSet struct {
	ctrl     *gomock.Controller
	recorder *MockDataSetMockRecorder
}

// MockDataSetMockRecorder is the mock recorder for MockDataSet
type MockDataSetMockRecorder struct {
	mock *MockDataSet
}

// NewMockDataSet creates a new mock instance
func NewMockDataSet(ctrl *gomock.Controller) *MockDataSet {
	mock := &MockDataSet{ctrl: ctrl}
	mock.recorder = &MockDataSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataSet) EXPECT() *MockDataSetMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockDataSet) Find(id uint8, name string) *Data {
	ret := m.ctrl.Call(m, "Find", id, name)
	ret0, _ := ret[0].(*Data)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockDataSetMockRecorder) Find(id, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDataSet)(nil).Find), id, name)
}

// Add mocks base method
func (m *MockDataSet) Add(data *Data) *dataSet {
	ret := m.ctrl.Call(m, "Add", data)
	ret0, _ := ret[0].(*dataSet)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockDataSetMockRecorder) Add(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDataSet)(nil).Add), data)
}

// Delete mocks base method
func (m *MockDataSet) Delete(id uint8, name string) *dataSet {
	ret := m.ctrl.Call(m, "Delete", id, name)
	ret0, _ := ret[0].(*dataSet)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDataSetMockRecorder) Delete(id, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataSet)(nil).Delete), id, name)
}
