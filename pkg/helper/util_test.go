package helper

import "testing"

func TestPathExists(t *testing.T) {
	exists, err := PathExists("../../conf/app.ini")
	if err != nil || !exists {
		t.Fail()
	}
}
