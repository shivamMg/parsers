package main

import (
	"errors"

	"github.com/alecthomas/chroma"
	"github.com/shivamMg/rd"
)

var lexer = chroma.MustNewLexer(
	&chroma.Config{
		Name: "Arithmetic Expressions",
	},
	chroma.Rules{
		"root": {
			{`\s+`, chroma.Text, nil},
			{`[+\-*/]`, chroma.Operator, nil},
			{`[()]`, chroma.Punctuation, nil},
			{`\d*\.\d+`, chroma.NumberFloat, nil},
			{`\d+`, chroma.NumberInteger, nil},
		},
	},
)

func Lex(expr string) (tokens []rd.Token, err error) {
	iter, err := lexer.Tokenise(nil, expr)
	if err != nil {
		return nil, err
	}
	for _, token := range iter.Tokens() {
		switch token.Type {
		case chroma.Operator, chroma.Punctuation, chroma.NumberInteger, chroma.NumberFloat:
			tokens = append(tokens, token.Value)
		case chroma.Error:
			return nil, errors.New("invalid token")
		}
	}
	return tokens, nil
}
