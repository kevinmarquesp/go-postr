package utils

import "testing"

func FailTestNowIfErrorIsNotNil(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
