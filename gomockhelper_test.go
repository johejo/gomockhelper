package gomockhelper_test

//go:generate mockgen -source=$GOFILE -destination=./internal/mock.go

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/johejo/gomockhelper"
	mock "github.com/johejo/gomockhelper/internal"
)

type Doer1 interface {
	Do1()
	Do2()
	Do3()
}

type Doer2 interface {
	DoWithReturns1(a int, b string) (string, int)
	DoWithReturns2(a int, b string) (string, int)
}

type Doer3 interface {
	Do1(a int, b string) (string, int)
	Do2(c bool, d []byte)
}

func TestExpectAll(t *testing.T) {

	t.Run("Doer1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var mockDoer1 Doer1 = mock.NewMockDoer1(ctrl)
		gomockhelper.ExpectAll(ctrl, mockDoer1, []interface{}{}, nil)

		mockDoer1.Do1()
		mockDoer1.Do2()
		mockDoer1.Do3()
	})

	t.Run("Doer2: with return value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var mockDoer2 Doer2 = mock.NewMockDoer2(ctrl)
		gomockhelper.ExpectAll(ctrl, mockDoer2, []interface{}{1234, "arg"}, func(call *gomock.Call) {
			call.Return("foo", 5050)
		})

		{
			got1, got2 := mockDoer2.DoWithReturns1(1234, "arg")
			if got1 != "foo" {
				t.Errorf("should return foo, but got=%v", got1)
			}
			if got2 != 5050 {
				t.Errorf("should return 5050, but got %v", got2)
			}
		}
		{
			got1, got2 := mockDoer2.DoWithReturns2(1234, "arg")
			if got1 != "foo" {
				t.Errorf("should return foo, but got=%v", got1)
			}
			if got2 != 5050 {
				t.Errorf("should return 5050, but got %v", got2)
			}
		}
	})

	t.Run("Doer3: with gomock.Any()", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var mockDoer3 Doer3 = mock.NewMockDoer3(ctrl)
		gomockhelper.ExpectAll(ctrl, mockDoer3, nil, nil)
		mockDoer3.Do1(1234, "arg")
		mockDoer3.Do2(true, []byte("aaa"))
	})
}

func Test_generate(t *testing.T) {
	if ci, _ := strconv.ParseBool(os.Getenv("CI")); !ci {
		t.Skip()
	}

	if err := exec.Command("go", "install", "github.com/golang/mock/mockgen").Run(); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command("go", "generate", ".").Run(); err != nil {
		t.Fatal(err)
	}
	out, err := exec.Command("git", "diff", ".").Output()
	if err != nil {
		t.Fatal(err)
	}
	diff := strings.TrimSpace(string(out))
	if diff != "" {
		t.Error("diff exists")
		t.Log(diff)
	}
}
