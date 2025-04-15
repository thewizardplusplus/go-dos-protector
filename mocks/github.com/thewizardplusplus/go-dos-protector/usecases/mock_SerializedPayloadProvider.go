// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package dosProtectorUsecasesMocks

import (
	"context"

	mock "github.com/stretchr/testify/mock"
	powValueTypes "github.com/thewizardplusplus/go-pow/value-types"
)

// NewMockSerializedPayloadProvider creates a new instance of MockSerializedPayloadProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSerializedPayloadProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSerializedPayloadProvider {
	mock := &MockSerializedPayloadProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockSerializedPayloadProvider is an autogenerated mock type for the SerializedPayloadProvider type
type MockSerializedPayloadProvider struct {
	mock.Mock
}

type MockSerializedPayloadProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSerializedPayloadProvider) EXPECT() *MockSerializedPayloadProvider_Expecter {
	return &MockSerializedPayloadProvider_Expecter{mock: &_m.Mock}
}

// ProvideSerializedPayload provides a mock function for the type MockSerializedPayloadProvider
func (_mock *MockSerializedPayloadProvider) ProvideSerializedPayload(ctx context.Context) (powValueTypes.SerializedPayload, error) {
	ret := _mock.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ProvideSerializedPayload")
	}

	var r0 powValueTypes.SerializedPayload
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context) (powValueTypes.SerializedPayload, error)); ok {
		return returnFunc(ctx)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context) powValueTypes.SerializedPayload); ok {
		r0 = returnFunc(ctx)
	} else {
		r0 = ret.Get(0).(powValueTypes.SerializedPayload)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = returnFunc(ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockSerializedPayloadProvider_ProvideSerializedPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProvideSerializedPayload'
type MockSerializedPayloadProvider_ProvideSerializedPayload_Call struct {
	*mock.Call
}

// ProvideSerializedPayload is a helper method to define mock.On call
//   - ctx
func (_e *MockSerializedPayloadProvider_Expecter) ProvideSerializedPayload(ctx interface{}) *MockSerializedPayloadProvider_ProvideSerializedPayload_Call {
	return &MockSerializedPayloadProvider_ProvideSerializedPayload_Call{Call: _e.mock.On("ProvideSerializedPayload", ctx)}
}

func (_c *MockSerializedPayloadProvider_ProvideSerializedPayload_Call) Run(run func(ctx context.Context)) *MockSerializedPayloadProvider_ProvideSerializedPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSerializedPayloadProvider_ProvideSerializedPayload_Call) Return(serializedPayload powValueTypes.SerializedPayload, err error) *MockSerializedPayloadProvider_ProvideSerializedPayload_Call {
	_c.Call.Return(serializedPayload, err)
	return _c
}

func (_c *MockSerializedPayloadProvider_ProvideSerializedPayload_Call) RunAndReturn(run func(ctx context.Context) (powValueTypes.SerializedPayload, error)) *MockSerializedPayloadProvider_ProvideSerializedPayload_Call {
	_c.Call.Return(run)
	return _c
}
