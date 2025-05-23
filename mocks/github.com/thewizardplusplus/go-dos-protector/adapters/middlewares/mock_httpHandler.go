// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package dosProtectorAdapterMiddlewaresMocks

import (
	"net/http"

	mock "github.com/stretchr/testify/mock"
)

// NewMockhttpHandler creates a new instance of MockhttpHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockhttpHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockhttpHandler {
	mock := &MockhttpHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockhttpHandler is an autogenerated mock type for the httpHandler type
type MockhttpHandler struct {
	mock.Mock
}

type MockhttpHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockhttpHandler) EXPECT() *MockhttpHandler_Expecter {
	return &MockhttpHandler_Expecter{mock: &_m.Mock}
}

// ServeHTTP provides a mock function for the type MockhttpHandler
func (_mock *MockhttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	_mock.Called(responseWriter, request)
	return
}

// MockhttpHandler_ServeHTTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServeHTTP'
type MockhttpHandler_ServeHTTP_Call struct {
	*mock.Call
}

// ServeHTTP is a helper method to define mock.On call
//   - responseWriter
//   - request
func (_e *MockhttpHandler_Expecter) ServeHTTP(responseWriter interface{}, request interface{}) *MockhttpHandler_ServeHTTP_Call {
	return &MockhttpHandler_ServeHTTP_Call{Call: _e.mock.On("ServeHTTP", responseWriter, request)}
}

func (_c *MockhttpHandler_ServeHTTP_Call) Run(run func(responseWriter http.ResponseWriter, request *http.Request)) *MockhttpHandler_ServeHTTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockhttpHandler_ServeHTTP_Call) Return() *MockhttpHandler_ServeHTTP_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockhttpHandler_ServeHTTP_Call) RunAndReturn(run func(responseWriter http.ResponseWriter, request *http.Request)) *MockhttpHandler_ServeHTTP_Call {
	_c.Run(run)
	return _c
}
