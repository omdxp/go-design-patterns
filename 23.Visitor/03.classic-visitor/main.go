package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(v ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(v ExpressionVisitor) {
	v.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(v ExpressionVisitor) {
	v.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
	return e.sb.String()
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteRune('(')
	ae.left.Accept(e)
	e.sb.WriteRune('+')
	ae.right.Accept(e)
	e.sb.WriteRune(')')
}

type ExpressionEvaluator struct {
	result float64
}

func NewExpressionEvaluator() *ExpressionEvaluator {
	return &ExpressionEvaluator{}
}

func (e *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	e.result = de.value
}

func (e *ExpressionEvaluator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(e)
	a := e.result
	ae.right.Accept(e)
	e.result += a
}

func main() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			&DoubleExpression{2},
			&DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	ep.VisitAdditionExpression(e)

	ee := NewExpressionEvaluator()
	ee.VisitAdditionExpression(e)
	fmt.Printf("%s = %g\n", ep, ee.result)
}

// Classic Visitor (double dispatch)
