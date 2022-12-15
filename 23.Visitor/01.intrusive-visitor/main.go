package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

func (a *AdditionExpression) String() string {
	sb := strings.Builder{}
	a.Print(&sb)
	return sb.String()
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

// Intrusive Visitor
