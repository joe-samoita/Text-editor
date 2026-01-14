package pipeline

import (

	"go-reloaded/internal/token"
	"go-reloaded/internal/rules"

)

type TestFunc func ([]token.Token) []token.Token 

func Run(tokens[]token.Token) []token.Token {
	
	Tests := []TestFunc{

		rules.Hexabin,
		rules.Uplowcap,
		rules.Punct,
		rules.Quote,
		rules.Articles,


	}

	for _, test := range Tests {
		tokens := test(tokens)
	}

	return tokens

}
