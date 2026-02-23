package main

import (
	"os"
	"strconv"
	"testing"

	hcl "github.com/openTaho/taho-go-hcl"
	thinHclParser "github.com/openTaho/taho-go-hcl-thin"
)

func testNewFile01(t *testing.T, parser hcl.HclParser) {
	file := parser.NewFile("./tests/01/01.tfvars")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	expect(t, file.Name(), wd+"/tests/01/01.tfvars")

	elements := file.Elements()
	expect(t, strconv.Itoa(len(elements)), "5")
	for i := 1; i < 4; i++ {
		el := elements[i]
		expect(t, el.Pair().Value(), "test"+strconv.Itoa(i))
		expect(t, el.Value(), "null")
	}
}

func expect(t *testing.T, value string, expecting string) {
	if value != expecting {
		t.Errorf("got \"%s\"; expected \"%s\"", value, expecting)
	}
}

func testNewDir(t *testing.T, parser hcl.HclParser) {
	dir := parser.NewDir("./tests/01")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	expect(t, dir.Name(), wd+"/tests/01")
	files, err := dir.Files()
	if err != nil {
		panic(err)
	}
	expect(t, strconv.Itoa(len(files)), "4")
	expect(t, files[0].Name(), wd+"/tests/01/01.tfvars")
	expect(t, files[1].Name(), wd+"/tests/01/02.tf")
	expect(t, files[2].Name(), wd+"/tests/01/03.hcl")
	expect(t, files[3].Name(), wd+"/tests/01/04.tf")
}

func TestThinNewDir(t *testing.T) {
	testNewDir(t, thinHclParser.New())
}

func TestThinNewFile01(t *testing.T) {
	testNewFile01(t, thinHclParser.New())
}
