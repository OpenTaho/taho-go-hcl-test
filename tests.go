package main

import (
	"os"
	"strconv"

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
	expect(t, strconv.Itoa(len(files)), "9")
	expect(t, files[0].Name(), wd+"/tests/01/000.tfvars")
	expect(t, files[1].Name(), wd+"/tests/01/001.tfvars")
	expect(t, files[2].Name(), wd+"/tests/01/002.tfvars")
	expect(t, files[3].Name(), wd+"/tests/01/003.tfvars")
	expect(t, files[4].Name(), wd+"/tests/01/004.tfvars")
	expect(t, files[5].Name(), wd+"/tests/01/101.tfvars")
	expect(t, files[6].Name(), wd+"/tests/01/102.tf")
	expect(t, files[7].Name(), wd+"/tests/01/103.hcl")
	expect(t, files[8].Name(), wd+"/tests/01/104.tf")
}

func testNewFile000(t *testing.T, parser hcl.HclParser) {
	file := parser.NewFile("./tests/01/000.tfvars")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	expect(t, file.Name(), wd+"/tests/01/000.tfvars")
	nodes := file.Nodes()
	expect(t, nodes[0].Value(), "000.tfvars")
}

func testNewFile101(t *testing.T, parser hcl.HclParser) {
	file := parser.NewFile("./tests/01/101.tfvars")

	nodes := file.Nodes()
	nLen := len(nodes)
	expect(t, strconv.Itoa(nLen), "1")
	expect(t, nodes[0].Value(), "101.tfvars")
	nodes = nodes[0].Nodes()

	// TODO: add this back later
	// nLen = len(nodes)
	// expect(t, strconv.Itoa(nLen), "17")

	// Build the string we need for our expectation
	txt := NewTextBuilder()
	txt.Add("# This sample content has a variety of elements, including comments, blank")
	txt.Add("# lines, null values, and multi-line strings. The parser should be able to")
	txt.Add("# handle all of these elements correctly and produce the expected output.")
	txt.Add("#")
	txt.Add("# This file should have 17 elements.")
	txt.Add("#")
	txt.Add("# [0]    - the file block")
	txt.Add("# [0][0] - the multi-line comment")
	txt.Add("# [0][1] - a pair with a null")
	txt.Add("# [0][2] - a pair with a null without whitespace")
	txt.Add("# [0][3] - a pair with a null with whitespace on the right")
	txt.Add("# [0][4] - a blank line")
	txt.Add("# [0][5] - a pair with a string")
	txt.Add("# [0][6] - a pair with a string without whitespace")
	txt.Add("# [0][7] - a pair with a string whitespace on the right")
	txt.Add("# [0][8] - another blank line")
	txt.Add("# [0][9] - a pair with a long name and another null value")
	txt.Add("# [0][10] - yet another blank line")
	txt.Add("# [0][11] - a pair with a heredoc")
	txt.Add("# [0][10] - yet another blank line")
	txt.Add("# [0][11] - a pair with an indent style heredoc")
	txt.Add("# [0][12] - a pair with a numeric")
	txt.Add("# [0][13] - a pair with a numeric without whitespace")
	txt.Add("# [0][14] - a pair with a numeric with whitespace on the right")
	txt.Add("# [0][15] - a pair with a numeric with whitespace on the left")
	txt.Add("#")
	txt.Add("")

	n := 0
	node := nodes[n]
	expect(t, node.Type().String(), "comment")
	expect(t, node.Value(), txt.String())

	n++
	node = nodes[n]
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Pair().Value(), "test1")
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "null")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test2")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "null")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test3")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "null")

	n++
	node = nodes[n]
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "\n")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test4")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "string")
	expect(t, node.Value(), "A \\\"simple\\\" string")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test5")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "string")
	expect(t, node.Value(), "A \\\"simple\\\" string")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test6")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "string")
	expect(t, node.Value(), "A \\\"simple\\\" string")

	n++
	node = nodes[n]
	expect(t, node.Value(), "\n")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test7_long_name")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "null")

	n++
	node = nodes[n]
	expect(t, node.Value(), "\n")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test8")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "doc")
	expect(t, node.Value(), "This is a heredoc.\n")

	n++
	node = nodes[n]
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "\n")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test9")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Type().String(), "doc-with-indent")
	expect(t, node.Value(), "This is an indent style heredoc.\n")

	n++
	node = nodes[n]
	expect(t, node.Type().String(), "token")
	expect(t, node.Value(), "\n")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test10")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Value(), "1")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test11")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Value(), "2")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test12")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Value(), "3")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test13")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Value(), "4")

	n++
	node = nodes[n]
	expect(t, node.Pair().Value(), "test14")
	expect(t, node.Pair().Type().String(), "token")
	expect(t, node.Value(), "5")
}
