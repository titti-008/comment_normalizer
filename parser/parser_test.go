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
			opts:    &Options{Symbol: SYMBOL_HASH},
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
			opts:    &Options{Symbol: SYMBOL_HASH},
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
			opts:    &Options{Symbol: SYMBOL_SLASH},
		},
		{
			name:    "Separated by CRLF",
			input:   "\r\n// This\r\n// is\r\n// a\r\n// comment.\r\n",
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{Newline: CRLF},
		},
		{
			name:    "Separated by CR",
			input:   "\r// This\r// is\r// a\r// comment.\r",
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{Newline: CR},
		},
		{
			name: "many tab in front of comment",
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
		{
			name: "many space in front of comment",
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
		{
			name:    "many space in front of comment with CRLF and hash symbol",
			input:   "\t\t\t# This\r\n\t\t\t# is\r\n\t\t\t# a\r\n\t\t\t# comment.\r\n",
			want:    `This is a comment.`,
			wantErr: false,
			opts:    &Options{Newline: CRLF, Symbol: SYMBOL_HASH},
		},
		{
			name: "many line comment has blank line in between",
			input: `
                    // This
                    // is
					// first
                    // comment.
					//
                    // This
                    // is
					// second
                    // comment.
					//
                    // This
                    // is
					// third
                    // comment.
                    `,
			want:    "This is first comment.\n\nThis is second comment.\n\nThis is third comment.",
			wantErr: false,
			opts:    &Options{},
		},
		{
			name: "many line comment has many blank line in between",
			input: `
                    // This
                    // is
					// first
                    // comment.
					//
					//
                    // This
                    // is
					// second
                    // comment.
					//
					//
					//
                    // This
                    // is
					// third
                    // comment.
                    `,
			want:    "This is first comment.\n\nThis is second comment.\n\nThis is third comment.",
			wantErr: false,
			opts:    &Options{},
		},
		{
			name: "Join sentence specified number of blank line when many line comment has many blank line in between",
			input: `
                    // This
                    // is
					// first
                    // comment.
					//
					//
                    // This
                    // is
					// second
                    // comment.
					//
					//
					//
                    // This
                    // is
					// third
                    // comment.
                    `,
			want:    "This is first comment.\n\n\n\nThis is second comment.\n\n\n\nThis is third comment.",
			wantErr: false,
			opts:    &Options{Join: 3},
		},
		{
			name:    "Join sentence specified number of blank line when many line comment has many blank line in between with CRLF",
			input:   " \t  // This\r\n// is\r\n// first\r\n// comment.\r\n\r\nAnd second line.",
			want:    "This is first comment.\r\n\r\nAnd second line.",
			wantErr: false,
			opts:    &Options{Join: 1, Newline: CRLF},
		},

		// TODO: コメント以外の行を無視する
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
