package parser

import (
	"strings"
)

const (
	// newline is the default newline character.
	_    newline = iota
	CR           // "\r"
	CRLF         // "\r\n"
	LF           // "\n"
)

const (
	// symbol is the default comment symbol.
	SYMBOL_HASH    = "#"
	SYMBOL_SLASH   = "//"
	SYMBOL_DEFAULT = SYMBOL_SLASH
)

type newline int
type symbol string

type Parser struct {
	input   string
	options *Options
}

type Options struct {
	newline newline
	symbol  symbol
}

func New(input string, opts *Options) *Parser {
	if opts.symbol == "" {
		opts.symbol = SYMBOL_DEFAULT
	}
	return &Parser{
		input:   input,
		options: opts,
	}
}

func (p *Parser) Parse() (string, error) {
	target := string(p.options.symbol)
	result := strings.Replace(p.input, target, "", -1)
	result = strings.TrimSpace(result)
	result = p.replaceNewLine(result)

	return result, nil
}

func (p *Parser) replaceNewLine(input string) string {
	switch p.options.newline {
	case CR:
		return strings.Replace(input, "\r", "", -1)
	case CRLF:
		return strings.Replace(input, "\r\n", "", -1)
	case LF:
		return strings.Replace(input, "\n", "", -1)
	default:
		return strings.Replace(input, "\n", "", -1)
	}

}
