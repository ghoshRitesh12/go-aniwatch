package goaniwatch_test

import (
	"fmt"
	"testing"

	goaniwatch "github.com/ghoshRitesh12/go-aniwatch"
	"github.com/ghoshRitesh12/go-aniwatch/parser"
)

func TestNewAniwatch(t *testing.T) {
	aniwatch := goaniwatch.New(parser.HeaderArgs{
		ACCEPT_HEADER: "something",
	})
	fmt.Printf("%+v", aniwatch)
	aniwatch.GetParse()
}
