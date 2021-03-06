// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/metalscreame/GoToBoox/src/dataBase/repository (interfaces: CommentsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	repository "github.com/metalscreame/GoToBoox/src/dataBase/repository"
	reflect "reflect"
)

// MockCommentsRepository is a mock of CommentsRepository interface
type MockCommentsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentsRepositoryMockRecorder
}

// MockCommentsRepositoryMockRecorder is the mock recorder for MockCommentsRepository
type MockCommentsRepositoryMockRecorder struct {
	mock *MockCommentsRepository
}

// NewMockCommentsRepository creates a new mock instance
func NewMockCommentsRepository(ctrl *gomock.Controller) *MockCommentsRepository {
	mock := &MockCommentsRepository{ctrl: ctrl}
	mock.recorder = &MockCommentsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentsRepository) EXPECT() *MockCommentsRepositoryMockRecorder {
	return m.recorder
}

// GetAllCommentsByBookID mocks base method
func (m *MockCommentsRepository) GetAllCommentsByBookID(arg0 int) ([]repository.Comment, error) {
	ret := m.ctrl.Call(m, "GetAllCommentsByBookID", arg0)
	ret0, _ := ret[0].([]repository.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommentsByBookID indicates an expected call of GetAllCommentsByBookID
func (mr *MockCommentsRepositoryMockRecorder) GetAllCommentsByBookID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommentsByBookID", reflect.TypeOf((*MockCommentsRepository)(nil).GetAllCommentsByBookID), arg0)
}

// GetAllCommentsByNickname mocks base method
func (m *MockCommentsRepository) GetAllCommentsByNickname(arg0 string) ([]repository.Comment, error) {
	ret := m.ctrl.Call(m, "GetAllCommentsByNickname", arg0)
	ret0, _ := ret[0].([]repository.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommentsByNickname indicates an expected call of GetAllCommentsByNickname
func (mr *MockCommentsRepositoryMockRecorder) GetAllCommentsByNickname(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommentsByNickname", reflect.TypeOf((*MockCommentsRepository)(nil).GetAllCommentsByNickname), arg0)
}

// InsertNewComment mocks base method
func (m *MockCommentsRepository) InsertNewComment(arg0, arg1, arg2 string, arg3 int) error {
	ret := m.ctrl.Call(m, "InsertNewComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewComment indicates an expected call of InsertNewComment
func (mr *MockCommentsRepositoryMockRecorder) InsertNewComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewComment", reflect.TypeOf((*MockCommentsRepository)(nil).InsertNewComment), arg0, arg1, arg2, arg3)
}
