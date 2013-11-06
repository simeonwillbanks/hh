package hh

import (
	"github.com/codegangsta/cli"
	"github.com/simeonwillbanks/hh/app"
	"io/ioutil"
	"testing"
)

func TestAction(t *testing.T) {
	f, _ := ioutil.TempFile("", "")

	hh.ActionWriter = f

	var c cli.Context
	hh.Action(&c)

	contents, _ := ioutil.ReadFile(f.Name())

	if string(contents) != "Hello Healthy Hacker!\n" {
		t.Error("TestAction test failed: Wrong message")
	}
}
