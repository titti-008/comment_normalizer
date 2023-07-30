package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/titti-008/comment_normalizer/parser"
)

func main() {
	inFile := flag.String("f", "", "Input file path")
	symbol := flag.String("s", "//", "Comment symbol. Default: //")

	flag.Parse()

	var input []byte

	if *inFile == "" {
		panic("Input file path is required")
	} else {
		f, err := os.Open(*inFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v", err)
			return
		}
		defer f.Close()

		buf := bufio.NewReader(f)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				break
			}
			input = append(input, line...)
			input = append(input, '\n')
		}
	}

	p := parser.New(string(input), &parser.Options{Symbol: parser.Symbol(*symbol)})

	result, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
