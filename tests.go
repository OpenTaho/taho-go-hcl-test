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
	nodes := file.Body()
	expect(t, nodes[0].Value(), "000.tfvars")
}

func testNewFile101(t *testing.T, parser hcl.HclParser) {
	file := parser.NewFile("./tests/01/101.tfvars")

	fileNodes := file.Body()
	nLen := len(fileNodes)
	expect(t, strconv.Itoa(nLen), "1")
	expect(t, fileNodes[0].Value(), "101.tfvars")
	fileNodes = fileNodes[0].Body()

	// Build the string we need for our expectation
	txt := NewTextBuilder()
	txt.Add("# This sample content has a variety of elements, including comments, blank")
	txt.Add("# lines, null values, and multi-line strings. The parser should be able to")
	txt.Add("# handle all of these elements correctly and produce the expected output.")
	txt.Add("#")
	txt.Add("# [0] ----- the file block")
	txt.Add("# [0][0] -- the multi-line header comment")
	txt.Add("# [0][1] -- a blank line")
	txt.Add("# [0][2] -- test01; a pair with a complex string")
	txt.Add("# [0][3] -- test02; a pair with a simple string")
	txt.Add("# [0][4] -- test03; another pair with a simple string")
	txt.Add("# [0][5] -- another blank line")
	txt.Add("# [0][6] -- test04; a pair with a null")
	txt.Add("# [0][7] -- test06; another pair with a null")
	txt.Add("# [0][8] -- another blank line")
	txt.Add("# [0][9] -- A two line comment for test05")
	txt.Add("# [0][10] - test05; a pair with a null")
	txt.Add("# [0][11] - another blank line")
	txt.Add("# [0][12] - A one line comment for a_test07_long_name")
	txt.Add("# [0][13] - a_test07_long_name; a pair with a long name and a null value")
	txt.Add("# [0][14] - another blank line")
	txt.Add("# [0][15] - test08; a pair with a document")
	txt.Add("# [0][16] - another blank line")
	txt.Add("# [0][17] - test09; another pair with a document (indent style)")
	txt.Add("# [0][18] - another blank line")
	txt.Add("# [0][19] - test10; a pair with a numeric value")
	txt.Add("# [0][20] - test11; another pair with a numeric value")
	txt.Add("# [0][21] - test12; another pair with a numeric value")
	txt.Add("# [0][22] - test13; another pair with a numeric value")
	txt.Add("# [0][23] - test13; a pair with an expression value")
	txt.Add("#")
	txt.Add("# Since a blank line follows this comment the comment here is not associated")
	txt.Add("# with the first pair.")
	txt.Add("#")
	txt.Add("# Later, we will have some examples of comments that are associated with pairs.")
	txt.Add("")

	node := fileNodes[0]
	nodes := node.Body()
	expect(t, node.Type().String(), "comment")
	expect(t, node.Value(), txt.String())

	node = fileNodes[1]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[2]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test01")
	expect(t, nodes[0].Type().String(), "string")
	expect(t, nodes[0].Value(), "A \\\"complex\\\" ${1 + 1} string")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[3]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test02")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "string")
	expect(t, nodes[1].Value(), "A simple string")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[4]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test03")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "string")
	expect(t, nodes[1].Value(), "A simple string")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[5]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[6]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test04")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "null")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[7]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test06")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "null")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[8]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[9]
	nodes = node.Body()
	expect(t, node.Type().String(), "comment")
	txt = NewTextBuilder()
	txt.Add("# This is a multi-line comment that will be associated with the test05 pair. It")
	txt.Add("# should move with the pair when the content is reformatted.")
	txt.Add("")
	expect(t, node.Value(), txt.String())

	node = fileNodes[10]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test05")
	expect(t, nodes[0].Type().String(), "token")
	expect(t, nodes[0].Value(), "null")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[11]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[12]
	nodes = node.Body()
	expect(t, node.Type().String(), "comment")
	txt = NewTextBuilder()
	txt.Add("# This is a comment that will be associated with the a_test7_long_name pair.")
	txt.Add("")
	expect(t, node.Value(), txt.String())

	node = fileNodes[13]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "a_test07_long_name")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "null")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[14]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[15]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test08")
	expect(t, nodes[0].Type().String(), "doc")
	expect(t, nodes[0].Value(), "This is a document.\n")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[16]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[17]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test09")
	expect(t, nodes[0].Type().String(), "doc-with-indent")
	expect(t, nodes[0].Value(), "This is an indent style document.\n")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[18]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[19]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test10")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "1")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[20]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test11")
	expect(t, nodes[0].Type().String(), "token")
	expect(t, nodes[0].Value(), "2")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[21]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test12")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "3")
	expect(t, strconv.Itoa(len(nodes)), "2")

	node = fileNodes[22]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test13")
	expect(t, nodes[0].Type().String(), "token")
	expect(t, nodes[0].Value(), "4")
	expect(t, strconv.Itoa(len(nodes)), "1")

	node = fileNodes[23]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test14")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "5")
	expect(t, nodes[2].Type().String(), "space")
	expect(t, nodes[2].Value(), " ")
	expect(t, nodes[3].Type().String(), "token")
	expect(t, nodes[3].Value(), "+")
	expect(t, nodes[4].Type().String(), "space")
	expect(t, nodes[4].Value(), " ")
	expect(t, nodes[5].Type().String(), "token")
	expect(t, nodes[5].Value(), "6")
	expect(t, strconv.Itoa(len(nodes)), "6")

	node = fileNodes[24]
	nodes = node.Body()
	expect(t, node.Type().String(), "space")
	expect(t, node.Value(), "\n")

	node = fileNodes[25]
	nodes = node.Body()
	expect(t, node.Type().String(), "pair")
	expect(t, node.Value(), "test15")
	expect(t, nodes[0].Type().String(), "space")
	expect(t, nodes[0].Value(), " ")
	expect(t, nodes[1].Type().String(), "token")
	expect(t, nodes[1].Value(), "7")
	expect(t, nodes[2].Type().String(), "space")
	expect(t, nodes[2].Value(), " ")
	expect(t, nodes[3].Type().String(), "token")
	expect(t, nodes[3].Value(), "-")
	expect(t, nodes[4].Type().String(), "space")
	expect(t, nodes[4].Value(), " ")
	expect(t, nodes[5].Type().String(), "token")
	expect(t, nodes[5].Value(), "8")
	expect(t, nodes[6].Type().String(), "space")
	expect(t, nodes[6].Value(), "          ")
	expect(t, nodes[7].Type().String(), "comment")
	expect(t, nodes[7].Value(), "# And a comment at the end of a line\n")
	expect(t, strconv.Itoa(len(nodes)), "8")

	nLen = len(fileNodes)
	expect(t, strconv.Itoa(nLen), "26")
}
