// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package dosProtectorAdapterMiddlewaresMocks

import (
	"net/http"

	mock "github.com/stretchr/testify/mock"
)

// NewMockhttpErrorHandler creates a new instance of MockhttpErrorHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockhttpErrorHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockhttpErrorHandler {
	mock := &MockhttpErrorHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockhttpErrorHandler is an autogenerated mock type for the httpErrorHandler type
type MockhttpErrorHandler struct {
	mock.Mock
}

type MockhttpErrorHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockhttpErrorHandler) EXPECT() *MockhttpErrorHandler_Expecter {
	return &MockhttpErrorHandler_Expecter{mock: &_m.Mock}
}

// HandleHTTPError provides a mock function for the type MockhttpErrorHandler
func (_mock *MockhttpErrorHandler) HandleHTTPError(writer http.ResponseWriter, err string, statusCode int) {
	_mock.Called(writer, err, statusCode)
	return
}

// MockhttpErrorHandler_HandleHTTPError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleHTTPError'
type MockhttpErrorHandler_HandleHTTPError_Call struct {
	*mock.Call
}

// HandleHTTPError is a helper method to define mock.On call
//   - writer
//   - err
//   - statusCode
func (_e *MockhttpErrorHandler_Expecter) HandleHTTPError(writer interface{}, err interface{}, statusCode interface{}) *MockhttpErrorHandler_HandleHTTPError_Call {
	return &MockhttpErrorHandler_HandleHTTPError_Call{Call: _e.mock.On("HandleHTTPError", writer, err, statusCode)}
}

func (_c *MockhttpErrorHandler_HandleHTTPError_Call) Run(run func(writer http.ResponseWriter, err string, statusCode int)) *MockhttpErrorHandler_HandleHTTPError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockhttpErrorHandler_HandleHTTPError_Call) Return() *MockhttpErrorHandler_HandleHTTPError_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockhttpErrorHandler_HandleHTTPError_Call) RunAndReturn(run func(writer http.ResponseWriter, err string, statusCode int)) *MockhttpErrorHandler_HandleHTTPError_Call {
	_c.Run(run)
	return _c
}
