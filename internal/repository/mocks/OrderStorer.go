// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// OrderStorer is an autogenerated mock type for the OrderStorer type
type OrderStorer struct {
	mock.Mock
}

// GetListOfOrder provides a mock function with given fields: ctx, userId, role
func (_m *OrderStorer) GetListOfOrder(ctx context.Context, userId int, role string) ([]dto.Order, error) {
	ret := _m.Called(ctx, userId, role)

	if len(ret) == 0 {
		panic("no return value specified for GetListOfOrder")
	}

	var r0 []dto.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) ([]dto.Order, error)); ok {
		return rf(ctx, userId, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) []dto.Order); ok {
		r0 = rf(ctx, userId, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, userId, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByID provides a mock function with given fields: ctx, orderID, userId, role
func (_m *OrderStorer) GetOrderByID(ctx context.Context, orderID int, userId int, role string) (dto.Order, error) {
	ret := _m.Called(ctx, orderID, userId, role)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByID")
	}

	var r0 dto.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) (dto.Order, error)); ok {
		return rf(ctx, orderID, userId, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) dto.Order); ok {
		r0 = rf(ctx, orderID, userId, role)
	} else {
		r0 = ret.Get(0).(dto.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, orderID, userId, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderItemsByOrderID provides a mock function with given fields: ctx, orderID, role, userID
func (_m *OrderStorer) GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error) {
	ret := _m.Called(ctx, orderID, role, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderItemsByOrderID")
	}

	var r0 []dto.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, int) ([]dto.CartItem, error)); ok {
		return rf(ctx, orderID, role, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, int) []dto.CartItem); ok {
		r0 = rf(ctx, orderID, role, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.CartItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, int) error); ok {
		r1 = rf(ctx, orderID, role, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderStorer creates a new instance of OrderStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderStorer {
	mock := &OrderStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
