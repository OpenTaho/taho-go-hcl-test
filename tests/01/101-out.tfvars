# This sample content has a variety of elements, including comments, blank
# lines, null values, and multi-line strings. The parser should be able to
# handle all of these elements correctly and produce the expected output.
#
# Since a blank line follows this comment the comment here is not associated
# with the first pair.
#
# Later, we will have some examples of comments that are associated with pairs.

test01 = "A \"complex\" ${"test"} string"
test02 = "A simple string"
test03_x = "A simple string"
test04 = null
test06 = null
test10 = 1
test11 = 2
test12 = 3
test13 = 4
test14 = 5 + 6
test15 = 7 - 8
test16 = (9 / 10)
test19 = "%{if 1+1 == 2 }true%{else}false%{endif}"
test20 = [a, b, c]

# This is a comment that will be associated with the a_test7_long_name pair.
a_test07_long_name = null

# This is a multi-line comment that will be associated with the test05 pair. It
# should move with the pair when the content is reformatted.
test05 = null

test08 = <<-EOT1
This is a document.
EOT1

test09 = <<-EOT2
This is an indent style document.
EOT2

test17 = (
  11 * 12
)

test18 = "some text and ${<<-EOT3
    a document with three lines
    the second line
    the third line
  EOT3
}"

test21 = [
  a,
[
    b.2,
    b.1,
    b.3,
  ]  ,
  c,
]

test22 = {
  b = "b"
  a = "a"
  
c = [
    1, 2,
    3,
  ]

    d = "d"

  e = "e"
}
