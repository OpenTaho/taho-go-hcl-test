package main

import (
	"testing"

	thinHclParser "github.com/openTaho/taho-go-hcl-thin"
)

// TestThinNewDir tests the NewDir function of the HclParser interface.
//
// It creates a new HclDir and verfies that that we have the expected state.
// Since our Hcl parsing is primarily used to support Terraform processing
// content at the directory level makes sense because within the Terraform
// language syntax is evaluated based on the directory context.
func TestThinNewDir(t *testing.T) {
	testNewDir(t, thinHclParser.New())
}

// TestThinNewFile000 tests the NewFile function of the HclParser interface.
func TestThinNewFile000(t *testing.T) {
	testNewFile000(t, thinHclParser.New())
}

// TestThinNewFile101 tests the NewFile function of the HclParser interface.
func TestThinNewFile101(t *testing.T) {
	testNewFile101(t, thinHclParser.New())
}
