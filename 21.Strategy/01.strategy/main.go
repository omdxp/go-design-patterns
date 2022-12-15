package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(buffer *strings.Builder)
	End(buffer *strings.Builder)
	AddListItem(buffer *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) End(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString(fmt.Sprintf(" * %s\n", item))
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(buffer *strings.Builder) {
	buffer.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(buffer *strings.Builder) {
	buffer.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString(fmt.Sprintf("  <li>%s</li>\n", item))
}

type TextProcessor struct {
	buffer       *strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(outputFormat OutputFormat) *TextProcessor {
	tp := &TextProcessor{
		buffer: &strings.Builder{},
	}
	tp.SetOutputFormat(outputFormat)
	return tp
}

func (t *TextProcessor) SetOutputFormat(outputFormat OutputFormat) {
	switch outputFormat {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.listStrategy.Start(t.buffer)
	for _, item := range items {
		t.listStrategy.AddListItem(t.buffer, item)
	}
	t.listStrategy.End(t.buffer)
}

func (t *TextProcessor) Clear() {
	t.buffer.Reset()
}

func (t *TextProcessor) String() string {
	return t.buffer.String()
}

func main() {
	tp := NewTextProcessor(Markdown)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())

	tp.Clear()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())
}

// Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
