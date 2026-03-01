package main

import (
	"testing"

	thinHclParser "github.com/openTaho/taho-go-hcl-thin"
)

func TestThinNewDir(t *testing.T) {
	testNewDir(t, thinHclParser.New())
}

func TestThinNewFile000(t *testing.T) {
	testNewFile000(t, thinHclParser.New())
}

func TestThinNewFile101(t *testing.T) {
	testNewFile101(t, thinHclParser.New())
}
