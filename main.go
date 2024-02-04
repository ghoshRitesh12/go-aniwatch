package goaniwatch

import (
	"github.com/ghoshRitesh12/go-aniwatch/parser"
)

func New(args parser.HeaderArgs) *parser.Parser {
	parser := parser.InitParser(args)

	return parser
}
