package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func main() {
	input := "1+(2+3)"
	tokens := Lex(input)
	fmt.Println(tokens)
}

// Interpreter is a behavioral design pattern that lets you define a grammar for a language, and then create an interpreter that uses this grammar to interpret sentences in the language.

// Lexing is the process of converting a sequence of characters into a sequence of tokens.
// A token is a meaningful unit of text, such as a word, number, punctuation mark, or operator.
