# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# [0] ----- the file block
# [0][0] -- the multi-line header comment
# [0][1] -- a blank line
# [0][2] -- test01; a pair with a complex string
# [0][3] -- test02; a pair with a simple string
# [0][4] -- test03; another pair with a simple string
# [0][5] -- another blank line
# [0][6] -- test04; a pair with a null
# [0][7] -- test06; another pair with a null
# [0][8] -- another blank line
# [0][9] -- A two line comment for test05
# [0][10] - test05; a pair with a null
# [0][11] - another blank line
# [0][12] - A one line comment for a_test07_long_name
# [0][13] - a_test07_long_name; a pair with a long name and a null value
# [0][14] - another blank line
# [0][15] - test08; a pair with a document
# [0][16] - another blank line
# [0][17] - test09; another pair with a document (indent style)
# [0][18] - another blank line
# [0][19] - test10; a pair with a numeric value
# [0][20] - test11; another pair with a numeric value
# [0][21] - test12; another pair with a numeric value
# [0][22] - test13; another pair with a numeric value
# [0][23] - test13; a pair with an expression value
#
# Since a blank line follows this comment the comment here is not associated
# with the first pair.
#
# Later, we will have some examples of comments that are associated with pairs.

test01="A \"complex\" ${"test"} string"
test02 = "A simple string"
test03= "A simple string"

test04 = null
test06= null

# This is a multi-line comment that will be associated with the test05 pair. It
# should move with the pair when the content is reformatted.
test05=null

# This is a comment that will be associated with the a_test7_long_name pair.
a_test07_long_name = null

test08 = <<EOT1
This is a document.
EOT1

test09 = <<-EOT2
This is an indent style document.
EOT2

test10 = 1
test11=2
test12= 3
test13 =4          # And a comment at the end of a line
test14 = 5 + 6

test15 = 7 - 8
test16 = (9 / 10)

test17 = (
  11 * 12
)

test18 = "some text and ${<<-EOT3
    a document with three lines
    the second line
    the third line
  EOT3
}"

test19 = "%{if 1+1 == 2 }true%{else}false%{endif}"
