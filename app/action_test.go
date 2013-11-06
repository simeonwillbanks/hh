package hh

import (
	"bytes"
	"github.com/codegangsta/cli"
	"github.com/simeonwillbanks/hh/app"
	"testing"
)

func TestAction(t *testing.T) {
	buf := new(bytes.Buffer)
	hh.ActionWriter = buf

	var c cli.Context
	hh.Action(&c)

	if buf.String() != "Hello Healthy Hacker!\n" {
		t.Error("TestAction test failed: Wrong message")
	}
}
