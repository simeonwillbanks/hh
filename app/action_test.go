package hh

import (
	"github.com/codegangsta/cli"
	"github.com/simeonwillbanks/hh/app"
	"testing"
)

func TestOutput(t *testing.T) {
	var c cli.Context
	if hh.Action(&c) != "Hello Healthy Hacker!" {
		t.Error("TestOutput test failed.")
	}
}
