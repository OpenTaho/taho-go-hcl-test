package main

import (
	"os"
	"strconv"
	"strings"

	"testing"

	hcl "github.com/openTaho/taho-go-hcl"
)

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

func testNewFile01(t *testing.T, parser hcl.HclParser) {
	file := parser.NewFile("./tests/01/01.tfvars")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	expect(t, file.Name(), wd+"/tests/01/01.tfvars")

	elements := file.Elements()
	eLen := len(elements)
	expect(t, strconv.Itoa(eLen), "1")
	elements = elements[0].NestedElements()
	eLen = len(elements)
	expect(t, strconv.Itoa(eLen), "17")

	// Build the string we need for our expectation
	b := []string{}
	b = append(b, "# This sample content has a variety of elements, including comments, blank")
	b = append(b, "# lines, null values, and multi-line strings. The parser should be able to")
	b = append(b, "# handle all of these elements correctly and produce the expected output.")
	b = append(b, "#")
	b = append(b, "# This file should have 17 elements.")
	b = append(b, "#")
	b = append(b, "# [0]    - the file block")
	b = append(b, "# [0][0] - the multi-line comment")
	b = append(b, "# [0][1] - a pair with a null")
	b = append(b, "# [0][2] - a pair with a null without whitespace")
	b = append(b, "# [0][3] - a pair with a null with whitespace on the right")
	b = append(b, "# [0][4] - a blank line")
	b = append(b, "# [0][5] - a pair with a string")
	b = append(b, "# [0][6] - a pair with a string without whitespace")
	b = append(b, "# [0][7] - a pair with a string whitespace on the right")
	b = append(b, "# [0][8] - another blank line")
	b = append(b, "# [0][9] - a pair with a long name and another null value")
	b = append(b, "# [0][10] - yet another blank line")
	b = append(b, "# [0][11] - a pair with a multi line string")
	b = append(b, "# [0][12] - a pair with a numeric")
	b = append(b, "# [0][13] - a pair with a numeric without whitespace")
	b = append(b, "# [0][14] - a pair with a numeric with whitespace on the right")
	b = append(b, "# [0][15] - a pair with a numeric with whitespace on the left")
	b = append(b, "#")
	b = append(b, "")

	el_num := 0
	el := elements[el_num]
	expect(t, el.Value(), strings.Join(b, "\n"))

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test1")
	expect(t, el.Value(), "null")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test2")
	expect(t, el.Value(), "null")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test3")
	expect(t, el.Value(), "null")

	el_num++
	el = elements[el_num]
	expect(t, el.Value(), "\n")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test4")
	expect(t, el.Value(), "A \\\"simple\\\" string")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test5")
	expect(t, el.Value(), "A \\\"simple\\\" string")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test6")
	expect(t, el.Value(), "A \\\"simple\\\" string")

	el_num++
	el = elements[el_num]
	expect(t, el.Value(), "\n")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test7_long_name")
	expect(t, el.Value(), "null")

	el_num++
	el = elements[el_num]
	expect(t, el.Value(), "\n")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test8")
	expect(t, el.Value(), "This is a multi-line string.\n")

	el_num++
	el = elements[el_num]
	expect(t, el.Value(), "\n")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test9")
	expect(t, el.Value(), "1")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test10")
	expect(t, el.Value(), "2")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test11")
	expect(t, el.Value(), "3")

	el_num++
	el = elements[el_num]
	expect(t, el.Pair().Value(), "test12")
	expect(t, el.Value(), "4")
}
