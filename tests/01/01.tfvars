# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# This file should have 8 elements.
#
# [0]    - the file block
# [0][0] - the multi-line comment
# [0][1] - a pair where a null value
# [0][2] - a blank line
# [0][3] - a pair where the value is a string
# [0][4] - another blank line
# [0][5] - a pair with a long name and another null value
# [0][6] - yet another blank line
# [0][7] - a pair with a multi line string
#
test1 = null

test2 = "A \"simple\" string"

test3_long_name = null

test4 = <<EOT
This is a multi-line string.
EOT
