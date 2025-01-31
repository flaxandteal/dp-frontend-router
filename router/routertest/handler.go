// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package routertest

import (
	"github.com/ONSdigital/dp-frontend-router/router"
	"net/http"
	"sync"
)

// Ensure, that HandlerMock does implement router.Handler.
// If this is not the case, regenerate this file with moq.
var _ router.Handler = &HandlerMock{}

// HandlerMock is a mock implementation of router.Handler.
//
//     func TestSomethingThatUsesHandler(t *testing.T) {
//
//         // make and configure a mocked router.Handler
//         mockedHandler := &HandlerMock{
//             ServeHTTPFunc: func(in1 http.ResponseWriter, in2 *http.Request)  {
// 	               panic("mock out the ServeHTTP method")
//             },
//         }
//
//         // use mockedHandler in code that requires router.Handler
//         // and then make assertions.
//
//     }
type HandlerMock struct {
	// ServeHTTPFunc mocks the ServeHTTP method.
	ServeHTTPFunc func(in1 http.ResponseWriter, in2 *http.Request)

	// calls tracks calls to the methods.
	calls struct {
		// ServeHTTP holds details about calls to the ServeHTTP method.
		ServeHTTP []struct {
			// In1 is the in1 argument value.
			In1 http.ResponseWriter
			// In2 is the in2 argument value.
			In2 *http.Request
		}
	}
	lockServeHTTP sync.RWMutex
}

// ServeHTTP calls ServeHTTPFunc.
func (mock *HandlerMock) ServeHTTP(in1 http.ResponseWriter, in2 *http.Request) {
	if mock.ServeHTTPFunc == nil {
		panic("HandlerMock.ServeHTTPFunc: method is nil but Handler.ServeHTTP was just called")
	}
	callInfo := struct {
		In1 http.ResponseWriter
		In2 *http.Request
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockServeHTTP.Lock()
	mock.calls.ServeHTTP = append(mock.calls.ServeHTTP, callInfo)
	mock.lockServeHTTP.Unlock()
	mock.ServeHTTPFunc(in1, in2)
}

// ServeHTTPCalls gets all the calls that were made to ServeHTTP.
// Check the length with:
//     len(mockedHandler.ServeHTTPCalls())
func (mock *HandlerMock) ServeHTTPCalls() []struct {
	In1 http.ResponseWriter
	In2 *http.Request
} {
	var calls []struct {
		In1 http.ResponseWriter
		In2 *http.Request
	}
	mock.lockServeHTTP.RLock()
	calls = mock.calls.ServeHTTP
	mock.lockServeHTTP.RUnlock()
	return calls
}
