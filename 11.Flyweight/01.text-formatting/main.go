package main

import (
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText, make([]bool, len(plainText))}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText, []*TextRange{}}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func main() {
	ft := NewFormattedText("This is a brave new world")
	ft.Capitalize(10, 15)
	println(ft.String())

	bft := NewBetterFormattedText("Make Algeria great again")
	bft.Range(13, 18).Capitalize = true
	println(bft.String())
}

// Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.
// The Flyweight pattern suggests that you store only part of the state of an object externally, and that you pass the rest of the state to the methods as parameters.
