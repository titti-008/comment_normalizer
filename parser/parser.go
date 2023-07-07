package parser

import (
	"strings"
	"unicode"
)

const (
	// newline is the default newline character.
	_ newline = iota
	CR
	CRLF
	LF
)

const (
	// symbol is the default comment symbol.
	SYMBOL_HASH    = "#"
	SYMBOL_SLASH   = "//"
	SYMBOL_DEFAULT = SYMBOL_SLASH
)

const (
	DEFOULT_JOIN = 1
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
	join    int
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
	if opts.join == 0 {
		opts.join = DEFOULT_JOIN
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
		line = strings.TrimLeftFunc(line, isSpaceOrTab)
		result += line + p.options.newline.String()
	}

	return result
}

func isSpaceOrTab(r rune) bool {
	return unicode.IsSpace(r) && (r == ' ' || r == '\t')
}

func (p *Parser) replaceCommentSymbol(input string) string {
	return strings.Replace(input, string(p.options.symbol), "", -1)
}

func (p *Parser) replaceNewLine(input string) string {
	lines := strings.Split(input, p.options.newline.String())

	result := []string{}
	sentence := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			result = append(result, strings.Join(sentence, " "))
			if len(sentence) > 0 {
				result = append(result, p.emptyLine())
			}
			sentence = []string{}
		} else {
			sentence = append(sentence, line)
		}
	}

	result = append(result, strings.Join(sentence, " "))
	return strings.Join(result, "")
}

func (p *Parser) emptyLine() string {
	result := p.options.newline.String()
	for i := 0; i < p.options.join; i++ {
		result += p.options.newline.String()
	}
	return result
}
