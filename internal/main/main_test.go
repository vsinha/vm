package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var b bytes.Buffer
	err := Main(&b)
	if err != nil {
		t.Errorf("doSomething() error: %v", err)
	}

	// The linter auto-adds 'fmt' to the list of imported packages, so we must use it somewhere. This should be able to be removed i
	fmt.Printf("damn you")
}

