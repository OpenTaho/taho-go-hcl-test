package main

import "strings"

type TextBuilder struct {
	strings []string
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{
		strings: []string{},
	}
}

func (b *TextBuilder) Add(s string) {
	b.strings = append(b.strings, s)
}

func (b *TextBuilder) String() string {
	return strings.Join(b.strings, "\n")
}
