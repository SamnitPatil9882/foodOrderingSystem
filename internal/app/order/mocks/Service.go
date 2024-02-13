// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddOrderItem provides a mock function with given fields: ctx, orderItemId, quantity
func (_m *Service) AddOrderItem(ctx context.Context, orderItemId int, quantity int) error {
	ret := _m.Called(ctx, orderItemId, quantity)

	if len(ret) == 0 {
		panic("no return value specified for AddOrderItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, orderItemId, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInvoice provides a mock function with given fields: ctx, invoiceInfo
func (_m *Service) CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error) {
	ret := _m.Called(ctx, invoiceInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateInvoice")
	}

	var r0 dto.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.InvoiceCreation) (dto.Invoice, error)); ok {
		return rf(ctx, invoiceInfo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.InvoiceCreation) dto.Invoice); ok {
		r0 = rf(ctx, invoiceInfo)
	} else {
		r0 = ret.Get(0).(dto.Invoice)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.InvoiceCreation) error); ok {
		r1 = rf(ctx, invoiceInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeliveryList provides a mock function with given fields: ctx, userID
func (_m *Service) GetDeliveryList(ctx context.Context, userID int) ([]dto.Delivery, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetDeliveryList")
	}

	var r0 []dto.Delivery
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]dto.Delivery, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []dto.Delivery); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Delivery)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInvoiceByOrderID provides a mock function with given fields: ctx, orderID, userID, role
func (_m *Service) GetInvoiceByOrderID(ctx context.Context, orderID int, userID int, role string) (dto.Invoice, error) {
	ret := _m.Called(ctx, orderID, userID, role)

	if len(ret) == 0 {
		panic("no return value specified for GetInvoiceByOrderID")
	}

	var r0 dto.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) (dto.Invoice, error)); ok {
		return rf(ctx, orderID, userID, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) dto.Invoice); ok {
		r0 = rf(ctx, orderID, userID, role)
	} else {
		r0 = ret.Get(0).(dto.Invoice)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, orderID, userID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListOfOrders provides a mock function with given fields: ctx, userID, role
func (_m *Service) GetListOfOrders(ctx context.Context, userID int, role string) ([]dto.Order, error) {
	ret := _m.Called(ctx, userID, role)

	if len(ret) == 0 {
		panic("no return value specified for GetListOfOrders")
	}

	var r0 []dto.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) ([]dto.Order, error)); ok {
		return rf(ctx, userID, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) []dto.Order); ok {
		r0 = rf(ctx, userID, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, userID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByID provides a mock function with given fields: ctx, orderID, userID, role
func (_m *Service) GetOrderByID(ctx context.Context, orderID int, userID int, role string) (dto.Order, error) {
	ret := _m.Called(ctx, orderID, userID, role)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByID")
	}

	var r0 dto.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) (dto.Order, error)); ok {
		return rf(ctx, orderID, userID, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) dto.Order); ok {
		r0 = rf(ctx, orderID, userID, role)
	} else {
		r0 = ret.Get(0).(dto.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, orderID, userID, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderItemsByOrderID provides a mock function with given fields: ctx, orderID, role, userID
func (_m *Service) GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error) {
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

// GetOrderList provides a mock function with given fields: ctx
func (_m *Service) GetOrderList(ctx context.Context) []dto.CartItem {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderList")
	}

	var r0 []dto.CartItem
	if rf, ok := ret.Get(0).(func(context.Context) []dto.CartItem); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.CartItem)
		}
	}

	return r0
}

// RemoveOrderItem provides a mock function with given fields: ctx, id
func (_m *Service) RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveOrderItem")
	}

	var r0 []dto.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]dto.CartItem, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []dto.CartItem); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.CartItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeliveryInfo provides a mock function with given fields: ctx, updateInfo
func (_m *Service) UpdateDeliveryInfo(ctx context.Context, updateInfo dto.DeliveryUpdateRequst) error {
	ret := _m.Called(ctx, updateInfo)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeliveryInfo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.DeliveryUpdateRequst) error); ok {
		r0 = rf(ctx, updateInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
