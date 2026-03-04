# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# [0] ----- the file block
# [0][0] -- the multi-line comment
# [0][1] -- a blank line
# [0][2] -- test1; a pair with a null
# [0][3] -- test2; a pair with a null without whitespace
# [0][4] -- test3; a pair with a null with whitespace on the right
# [0][5] -- a blank line
# [0][6] -- test4; a pair with a string
# [0][7] -- test5; a pair with a string without whitespace
# [0][8] -- test6; a pair with a string whitespace on the right
# [0][9] -- another blank line
# [0][10] - test7_...; a pair with a long name and another null value
# [0][11] - yet another blank line
# [0][12] - test8; a pair with a heredoc
# [0][13] - yet another blank line
# [0][14] - test9; a pair with an indent style heredoc
# [0][15] - yet another blank line
# [0][16] - test10; a pair with a numeric
# [0][17] - test11; a pair with a numeric without whitespace
# [0][18] - test12; a pair with a numeric with whitespace on the right
# [0][19] - test13; a pair with a numeric with whitespace on the left
# [0][20] - test14; a pair with a expression of "5 + 6"
#
# Since a blank line follows this comment the comment here is not associated
# with the first pair.
#
# Later, we will have some examples of comments that are associated with pairs.

test1 = null
test2=null
test3= null

test4 = "A \"simple\" string"
test5="A \"simple\" string"
test6= "A \"simple\" string"

a_test7_long_name = null

test8 = <<EOT
This is a heredoc.
EOT

test9 = <<-EOT
This is an indent style heredoc.
EOT

test10 = 1
test11=2
test12= 3
test13 =4
test14 = 5 + 6
