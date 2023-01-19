package integration

import (
	"testing"
)

func ErrorFailNow(t *testing.T, err error) {
	if err != nil {
		t.Errorf("err: " + err.Error())
		t.FailNow()
	}
}

func ErrorFail(t *testing.T, err error) {
	if err != nil {
		t.Errorf("err: " + err.Error())
		t.Fail()
	}
}
