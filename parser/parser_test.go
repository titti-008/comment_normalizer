package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		opts    *Options
		want    string
		wantErr bool
	}{
		{
			name:    "one newline comment with default symbol",
			input:   `// one line go comment.`,
			want:    "one line go comment.",
			wantErr: false,
			opts:    &Options{},
		},
		{
			name:    "one newline comment with hash symbol",
			input:   `# one line ruby comment.`,
			want:    "one line ruby comment.",
			wantErr: false,
			opts:    &Options{symbol: SYMBOL_HASH},
		},
		{
			name: "Separated single line comment with hash symbol",
			input: `
# This
# is
# a
# comment.
					`,
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{symbol: SYMBOL_HASH},
		},
		{
			name: "Separated single line comment with slash symbol",
			input: `
// This
// is
// a
// comment.
					`,
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{symbol: SYMBOL_SLASH},
		},
		{
			name:    "Separated by CRLF",
			input:   "\r\n// This\r\n// is\r\n// a\r\n// comment.\r\n",
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{newline: CRLF},
		},
		{
			name:    "Separated by CR",
			input:   "\r// This\r// is\r// a\r// comment.\r",
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{newline: CR},
		},
		{
			name: "many tab and space in front of comment",
			input: `
					// This
					// is
					// a
					// comment.
					`,
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{},
		},
		// 		{
		// 			name: "default newline",
		// 			input: `
		// 						# This
		// 						# is
		// 						# a
		// 						# comment.
		// 						#
		// 						# and second comment.
		// 					`,
		// 			want: `
		// 		This is a comment.
		// 		and second comment.`,
		// 			wantErr: false,
		// 			opts:    Options{symbol: SYMBOL_HASH},
		// 		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.input, tt.opts)
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("Parse() want: %q, but got: %q", tt.want, got)
			}
		})
	}
}
