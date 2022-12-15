package main

import (
	"fmt"
	"strings"
)

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) String() string {
	sb := strings.Builder{}
	Print(a, &sb)
	return sb.String()
}

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
}

func main() {
	// 1 + (2 + 3)
	e := AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			&DoubleExpression{2},
			&DoubleExpression{3},
		},
	}
	fmt.Println(e.String())
}

// Reflective Visitor
