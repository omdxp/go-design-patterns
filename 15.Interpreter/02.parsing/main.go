package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value}
}

func (i *Integer) Value() int {
	return i.value
}

type Operation int

const (
	Add Operation = iota
	Subtract
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func NewBinaryOperation(Type Operation, Left, Right Element) *BinaryOperation {
	return &BinaryOperation{Type, Left, Right}
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Add:
		return b.Left.Value() + b.Right.Value()
	case Subtract:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("Unknown operation")
	}
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Value)
}

func Lex(input string) []Token {
	var result []Token
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Plus, "+"})
		case '-':
			result = append(result, Token{Minus, "-"})
		case '(':
			result = append(result, Token{Lparen, "("})
		case ')':
			result = append(result, Token{Rparen, ")"})
		default:
			var sb strings.Builder
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{Int, sb.String()})
					i--
					break
				}
			}
		}
	}
	return result
}

func Parse(tokens []Token) Element {
	result := NewBinaryOperation(Add, NewInteger(0), NewInteger(0))
	haveLHS := false
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case Int:
			intValue := 0
			fmt.Sscanf(token.Value, "%d", &intValue)
			integer := NewInteger(intValue)
			if !haveLHS {
				result.Left = integer
				haveLHS = true
			} else {
				result.Right = integer
			}
		case Plus:
			result.Type = Add
		case Minus:
			result.Type = Subtract
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			// process subexpression w/o opening (
			subexpression := tokens[i+1 : j]
			element := Parse(subexpression)
			if !haveLHS {
				result.Left = element
				haveLHS = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return result
}

func main() {
	input := "(1+2)-(3+4)"
	tokens := Lex(input)
	fmt.Println(tokens)
	result := Parse(tokens)
	fmt.Println(result.Value())
}

// Parsing is the process of turning a sequence of tokens into a tree of expressions.
