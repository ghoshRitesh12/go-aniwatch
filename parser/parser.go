package parser

import "github.com/ghoshRitesh12/go-aniwatch/utils"

type src struct {
	BASE_URL   string
	AJAX_URL   string
	HOME_URL   string
	SEARCH_URL string
}

type header struct {
	ACCEPT          string
	USER_AGENT      string
	ACCEPT_ENCODING string
}

type Parser struct {
	header header
	src    src
}

func Init() *Parser {
	parser := &Parser{
		src: src{
			BASE_URL:   utils.SRC_BASE_URL,
			HOME_URL:   utils.SRC_HOME_URL,
			AJAX_URL:   utils.SRC_AJAX_URL,
			SEARCH_URL: utils.SRC_SEARCH_URL,
		},
		header: header{
			ACCEPT:          utils.ACCEPT_HEADER,
			USER_AGENT:      utils.USER_AGENT_HEADER,
			ACCEPT_ENCODING: utils.ACCEPT_ENCODING_HEADER,
		},
	}

	return parser
}
