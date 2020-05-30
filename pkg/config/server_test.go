package config

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	Server = (&server{}).Load("../../conf/app.ini").Init()
	fmt.Println(Server)
	if Server == nil {
		t.Fail()
	}
}