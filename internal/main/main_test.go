package main

import (
	"bytes"
	"testing"
)

func TestInit(t *testing.T) {
	var b bytes.Buffer
	err := Main(&b)
	if err != nil {
		t.Errorf("doSomething() error: %v", err)
	}
}
