package main

import "strings"

// The TextBuilder type is a simple utility for building multi-line text. It
// allows you to add strings and then retrieve the final concatenated result.
type TextBuilder struct {
	strings []string
}

// NewTextBuilder creates and returns a new instance of TextBuilder.
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{
		strings: []string{},
	}
}

// Add appends a new string to the TextBuilder's internal slice of strings.
func (b *TextBuilder) Add(s string) {
	b.strings = append(b.strings, s)
}

// String returns the concatenated result of all strings added to the
// TextBuilder, separated by newline characters.
func (b *TextBuilder) String() string {
	return strings.Join(b.strings, "\n")
}
