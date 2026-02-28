# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# This file should have 17 elements.
#
# [0]    - the file block
# [0][0] - the multi-line comment
# [0][1] - a pair with a null
# [0][2] - a pair with a null without whitespace
# [0][3] - a pair with a null with whitespace on the right
# [0][4] - a blank line
# [0][5] - a pair with a string
# [0][6] - a pair with a string without whitespace
# [0][7] - a pair with a string whitespace on the right
# [0][8] - another blank line
# [0][9] - a pair with a long name and another null value
# [0][10] - yet another blank line
# [0][11] - a pair with a multi line string
# [0][12] - a pair with a numeric
# [0][13] - a pair with a numeric without whitespace
# [0][14] - a pair with a numeric with whitespace on the right
# [0][15] - a pair with a numeric with whitespace on the left
#
test1 = null
test2=null
test3= null

test4 = "A \"simple\" string"
test5="A \"simple\" string"
test6= "A \"simple\" string"

test7_long_name = null

test8 = <<EOT
This is a multi-line string.
EOT

test9 = 1
test10=2
test11= 3
test12 =4
