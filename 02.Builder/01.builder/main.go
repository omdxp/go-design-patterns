package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.StringImpl(0)
}

func (e *HtmlElement) StringImpl(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, v := range e.elements {
		sb.WriteString(v.StringImpl(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{name: rootName}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	hello := "Hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	println(sb.String())

	words := []string{"Hello", "World"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	println(sb.String())

	// builder
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	println(b.String())

	// fluent builder
	b = NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").AddChildFluent("li", "world")
	println(b.String())
}

// Builder design pattern is a creational design pattern that allows you to separate the construction of a complex object from its representation.
// It is used to create a complex object using simple objects and using a step by step approach.
// A Builder class builds the final object step by step. This builder is independent of other objects.

// The Builder pattern is used to separate the construction of a complex object from its representation so that the same construction process can create different representations.
// A builder is responsible for defining the steps required to assemble a complex object, and it provides an API for executing these steps.
// A builder is also known as a director.

// Fluent Builder is a design pattern that allows you to create complex objects using simple objects and a step-by-step approach.
// A fluent builder is a builder that returns the same object after each action.
// This allows you to chain multiple actions in a single statement.

// HtmlBuilder is a builder that builds HTML elements.
