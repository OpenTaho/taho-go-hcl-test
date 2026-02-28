# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# This file should have 10 elements.
#
# [0]    - the file block
# [0][0] - the multi-line comment
# [0][1] - a pair with a null value
# [0][2] - a pair with a null and without whitespace around the equals sign
# [0][3] - a pair with a null and whitespace on right of equals sign
# [0][4] - a blank line
# [0][5] - a pair where the value is a string
# [0][6] - another blank line
# [0][7] - a pair with a long name and another null value
# [0][8] - yet another blank line
# [0][9] - a pair with a multi line string
#
test1 = null
test2=null
test3= null

test4 = "A \"simple\" string"

test5_long_name = null

test6 = <<EOT
This is a multi-line string.
EOT
