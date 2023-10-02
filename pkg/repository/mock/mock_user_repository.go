// Code generated by MockGen. DO NOT EDIT.
// Source: Teeverse/pkg/repository/interface (interfaces: UserRepository)

// Package mockrepository is a generated GoMock package.
package mockrepository

import (
	domain "Teeverse/pkg/domain"
	models "Teeverse/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(arg0 int, arg1 models.AddAddress, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), arg0, arg1, arg2)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), arg0, arg1)
}

// CheckIfFirstAddress mocks base method.
func (m *MockUserRepository) CheckIfFirstAddress(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfFirstAddress", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfFirstAddress indicates an expected call of CheckIfFirstAddress.
func (mr *MockUserRepositoryMockRecorder) CheckIfFirstAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfFirstAddress", reflect.TypeOf((*MockUserRepository)(nil).CheckIfFirstAddress), arg0)
}

// CheckUserAvailability mocks base method.
func (m *MockUserRepository) CheckUserAvailability(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAvailability", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUserAvailability indicates an expected call of CheckUserAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckUserAvailability(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAvailability), arg0)
}

// ClearCart mocks base method.
func (m *MockUserRepository) ClearCart(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearCart indicates an expected call of ClearCart.
func (mr *MockUserRepositoryMockRecorder) ClearCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCart", reflect.TypeOf((*MockUserRepository)(nil).ClearCart), arg0)
}

// EditEmail mocks base method.
func (m *MockUserRepository) EditEmail(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditEmail indicates an expected call of EditEmail.
func (mr *MockUserRepositoryMockRecorder) EditEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEmail", reflect.TypeOf((*MockUserRepository)(nil).EditEmail), arg0, arg1)
}

// EditName mocks base method.
func (m *MockUserRepository) EditName(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditName indicates an expected call of EditName.
func (mr *MockUserRepositoryMockRecorder) EditName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditName", reflect.TypeOf((*MockUserRepository)(nil).EditName), arg0, arg1)
}

// EditPhone mocks base method.
func (m *MockUserRepository) EditPhone(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditPhone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditPhone indicates an expected call of EditPhone.
func (mr *MockUserRepositoryMockRecorder) EditPhone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPhone", reflect.TypeOf((*MockUserRepository)(nil).EditPhone), arg0, arg1)
}

// EditUsername mocks base method.
func (m *MockUserRepository) EditUsername(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUsername", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUsername indicates an expected call of EditUsername.
func (mr *MockUserRepositoryMockRecorder) EditUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUsername", reflect.TypeOf((*MockUserRepository)(nil).EditUsername), arg0, arg1)
}

// FindCartQuantity mocks base method.
func (m *MockUserRepository) FindCartQuantity(arg0, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartQuantity", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartQuantity indicates an expected call of FindCartQuantity.
func (mr *MockUserRepositoryMockRecorder) FindCartQuantity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartQuantity", reflect.TypeOf((*MockUserRepository)(nil).FindCartQuantity), arg0, arg1)
}

// FindCategory mocks base method.
func (m *MockUserRepository) FindCategory(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategory", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategory indicates an expected call of FindCategory.
func (mr *MockUserRepositoryMockRecorder) FindCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategory", reflect.TypeOf((*MockUserRepository)(nil).FindCategory), arg0)
}

// FindIdFromPhone mocks base method.
func (m *MockUserRepository) FindIdFromPhone(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIdFromPhone", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIdFromPhone indicates an expected call of FindIdFromPhone.
func (mr *MockUserRepositoryMockRecorder) FindIdFromPhone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIdFromPhone", reflect.TypeOf((*MockUserRepository)(nil).FindIdFromPhone), arg0)
}

// FindPrice mocks base method.
func (m *MockUserRepository) FindPrice(arg0 int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrice", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrice indicates an expected call of FindPrice.
func (mr *MockUserRepositoryMockRecorder) FindPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrice", reflect.TypeOf((*MockUserRepository)(nil).FindPrice), arg0)
}

// FindProductNames mocks base method.
func (m *MockUserRepository) FindProductNames(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductNames", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductNames indicates an expected call of FindProductNames.
func (mr *MockUserRepositoryMockRecorder) FindProductNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductNames", reflect.TypeOf((*MockUserRepository)(nil).FindProductNames), arg0)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(arg0 models.UserLogin) (models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), arg0)
}

// FindUserIDByOrderID mocks base method.
func (m *MockUserRepository) FindUserIDByOrderID(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserIDByOrderID", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserIDByOrderID indicates an expected call of FindUserIDByOrderID.
func (mr *MockUserRepositoryMockRecorder) FindUserIDByOrderID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserIDByOrderID", reflect.TypeOf((*MockUserRepository)(nil).FindUserIDByOrderID), arg0)
}

// GetAddresses mocks base method.
func (m *MockUserRepository) GetAddresses(arg0 int) ([]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", arg0)
	ret0, _ := ret[0].([]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAddresses), arg0)
}

// GetCartID mocks base method.
func (m *MockUserRepository) GetCartID(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartID", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartID indicates an expected call of GetCartID.
func (mr *MockUserRepositoryMockRecorder) GetCartID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartID", reflect.TypeOf((*MockUserRepository)(nil).GetCartID), arg0)
}

// GetPassword mocks base method.
func (m *MockUserRepository) GetPassword(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserRepositoryMockRecorder) GetPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserRepository)(nil).GetPassword), arg0)
}

// GetProductsInCart mocks base method.
func (m *MockUserRepository) GetProductsInCart(arg0, arg1, arg2 int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsInCart", arg0, arg1, arg2)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsInCart indicates an expected call of GetProductsInCart.
func (mr *MockUserRepositoryMockRecorder) GetProductsInCart(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsInCart", reflect.TypeOf((*MockUserRepository)(nil).GetProductsInCart), arg0, arg1, arg2)
}

// GetUserDetails mocks base method.
func (m *MockUserRepository) GetUserDetails(arg0 int) (models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", arg0)
	ret0, _ := ret[0].(models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockUserRepositoryMockRecorder) GetUserDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockUserRepository)(nil).GetUserDetails), arg0)
}

// RemoveFromCart mocks base method.
func (m *MockUserRepository) RemoveFromCart(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockUserRepositoryMockRecorder) RemoveFromCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockUserRepository)(nil).RemoveFromCart), arg0, arg1)
}

// SignUp mocks base method.
func (m *MockUserRepository) SignUp(arg0 models.UserDetails) (models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0)
	ret0, _ := ret[0].(models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserRepositoryMockRecorder) SignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserRepository)(nil).SignUp), arg0)
}

// UpdateQuantityAdd mocks base method.
func (m *MockUserRepository) UpdateQuantityAdd(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityAdd indicates an expected call of UpdateQuantityAdd.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantityAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityAdd", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantityAdd), arg0, arg1)
}

// UpdateQuantityLess mocks base method.
func (m *MockUserRepository) UpdateQuantityLess(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityLess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityLess indicates an expected call of UpdateQuantityLess.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantityLess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityLess", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantityLess), arg0, arg1)
}

// UserBlockStatus mocks base method.
func (m *MockUserRepository) UserBlockStatus(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBlockStatus", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBlockStatus indicates an expected call of UserBlockStatus.
func (mr *MockUserRepositoryMockRecorder) UserBlockStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBlockStatus", reflect.TypeOf((*MockUserRepository)(nil).UserBlockStatus), arg0)
}
