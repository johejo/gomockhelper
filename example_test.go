package gomockhelper_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/johejo/gomockhelper"
	mock "github.com/johejo/gomockhelper/internal"
)

func ExampleExpectAll() {
	t := new(testing.T) // *testing.T

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var mockDoer1 Doer1 = mock.NewMockDoer1(ctrl)
	gomockhelper.ExpectAll(ctrl, mockDoer1, []interface{}{}, nil)

	mockDoer1.Do1()
	mockDoer1.Do2()
	mockDoer1.Do3()

	var mockDoer2 Doer2 = mock.NewMockDoer2(ctrl)
	gomockhelper.ExpectAll(ctrl, mockDoer2, []interface{}{1234, "arg"}, func(call *gomock.Call) {
		call.Return("foo", 5050)
	})

	fmt.Println(mockDoer2.DoWithReturns1(1234, "arg"))
	fmt.Println(mockDoer2.DoWithReturns2(1234, "arg"))

	// Output:
	// foo 5050
	// foo 5050

	var mockDoer3 Doer3 = mock.NewMockDoer3(ctrl)
	gomockhelper.ExpectAll(ctrl, mockDoer3, nil, nil)
	mockDoer3.Do1(1234, "arg")
	mockDoer3.Do2(true, []byte("aaa"))
}
