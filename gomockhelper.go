package gomockhelper

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

type function struct {
	m    *reflect.Method
	args []interface{}
}

// ExpectAll expects all methods of target mock with gomock.
// Passing nil in args will use gomock.Any() for all arguments.
// If you want to use Do or DoAndReturn of gomock.Call, you can use it by passing a function to callFunc.
// Example: func(call *gomock.Call) { call.Return("return", 5050)}
func ExpectAll(ctrl *gomock.Controller, mock interface{}, args []interface{}, callFunc func(call *gomock.Call)) {
	mockType := reflect.TypeOf(mock)
	numMethods := mockType.NumMethod()
	funcs := make([]*function, 0, numMethods)

	for i := 0; i < numMethods; i++ {
		if m := mockType.Method(i); m.Name != "EXPECT" {
			if args == nil {
				numArgs := m.Type.NumIn() - 1 // without receiver
				args = make([]interface{}, 0, numArgs)
				for i := 0; i < numArgs; i++ {
					args = append(args, gomock.Any())
				}
			}
			funcs = append(funcs, &function{m: &m, args: args})
		}
	}
	for _, f := range funcs {
		f := f
		call := ctrl.RecordCallWithMethodType(mock, f.m.Name, f.m.Type, f.args...)
		if callFunc != nil {
			callFunc(call)
		}
	}
}
