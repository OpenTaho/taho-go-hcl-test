package main

import (
	"testing"

	thinHclParser "github.com/openTaho/taho-go-hcl-thin"
)

func TestThinNewDir(t *testing.T) {
	testNewDir(t, thinHclParser.New())
}

func TestThinNewFile01(t *testing.T) {
	testNewFile01(t, thinHclParser.New())
}
