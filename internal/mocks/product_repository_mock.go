// Code generated by mockery v2.42.1. DO NOT EDIT.

package repomocks

import (
	domain "github.com/rulanugrh/lysithea/internal/entity/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// CountProduct provides a mock function with given fields:
func (_m *ProductRepository) CountProduct() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CountProduct")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountProductByCategoryID provides a mock function with given fields: categoryID
func (_m *ProductRepository) CountProductByCategoryID(categoryID uint) (int64, error) {
	ret := _m.Called(categoryID)

	if len(ret) == 0 {
		panic("no return value specified for CountProductByCategoryID")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int64, error)); ok {
		return rf(categoryID)
	}
	if rf, ok := ret.Get(0).(func(uint) int64); ok {
		r0 = rf(categoryID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: req
func (_m *ProductRepository) Create(req domain.ProductRequest) (*domain.Product, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.ProductRequest) (*domain.Product, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(domain.ProductRequest) *domain.Product); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.ProductRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: page, perPage
func (_m *ProductRepository) FindAll(page int, perPage int) (*[]domain.Product, error) {
	ret := _m.Called(page, perPage)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 *[]domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (*[]domain.Product, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) *[]domain.Product); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCategoryID provides a mock function with given fields: page, perPage, categoryID
func (_m *ProductRepository) FindByCategoryID(page int, perPage int, categoryID uint) (*[]domain.Product, error) {
	ret := _m.Called(page, perPage, categoryID)

	if len(ret) == 0 {
		panic("no return value specified for FindByCategoryID")
	}

	var r0 *[]domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, uint) (*[]domain.Product, error)); ok {
		return rf(page, perPage, categoryID)
	}
	if rf, ok := ret.Get(0).(func(int, int, uint) *[]domain.Product); ok {
		r0 = rf(page, perPage, categoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, uint) error); ok {
		r1 = rf(page, perPage, categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindID provides a mock function with given fields: id
func (_m *ProductRepository) FindID(id uint) (*domain.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindID")
	}

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, req
func (_m *ProductRepository) Update(id uint, req domain.Product) (*domain.Product, error) {
	ret := _m.Called(id, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, domain.Product) (*domain.Product, error)); ok {
		return rf(id, req)
	}
	if rf, ok := ret.Get(0).(func(uint, domain.Product) *domain.Product); ok {
		r0 = rf(id, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, domain.Product) error); ok {
		r1 = rf(id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}