package parser

import (
	"strings"
	"unicode"
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

func (n newline) String() string {
	switch n {
	case CR:
		return "\r"
	case CRLF:
		return "\r\n"
	case LF:
		return "\n"
	default:
		return "\n"
	}
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
	result := p.trimLeadingSpacesTabs(p.input)
	result = p.replaceCommentSymbol(result)
	result = strings.TrimSpace(result)
	result = p.replaceNewLine(result)

	return result, nil
}

func (p *Parser) trimLeadingSpacesTabs(input string) string {
	lines := strings.Split(input, p.options.newline.String())
	result := ""
	for _, line := range lines {
		line = strings.TrimLeftFunc(line, func(r rune) bool {
			return unicode.IsSpace(r) && (r == ' ' || r == '\t')
		})
		result += line + p.options.newline.String()
	}

	return result
}

func (p *Parser) replaceCommentSymbol(input string) string {
	return strings.Replace(input, string(p.options.symbol), "", -1)
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
