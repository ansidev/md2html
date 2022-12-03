//go:build !windows

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetFirstLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Normal single line string",
			args: args{
				s: "Line 1",
			},
			want: "Line 1",
		},
		{
			name: "Normal multi-line string",
			args: args{
				s: `Line 1
Line 2`,
			},
			want: "Line 1",
		},
		{
			name: "Single line markdown heading",
			args: args{
				s: "# Line 1",
			},
			want: "Line 1",
		},
		{
			name: "Multi-line markdown heading",
			args: args{
				s: `# Line 1
Line 2`,
			},
			want: "Line 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFirstLine(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_AppendIndentation(t *testing.T) {
	type args struct {
		s            string
		indentString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String indent",
			args: args{
				s: `<p>Paragraph 1</p>
<p>Paragraph 2</p>`,
				indentString: "  ",
			},
			want: `  <p>Paragraph 1</p>
  <p>Paragraph 2</p>`,
		},
		{
			name: "Tab indent",
			args: args{
				s: `<p>Paragraph 1</p>
<p>Paragraph 2</p>`,
				indentString: "	",
			},
			want: `	<p>Paragraph 1</p>
	<p>Paragraph 2</p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendIndentation(tt.args.s, tt.args.indentString)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestTrimBlankLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Paragraph with blank lines",
			args: args{
				s: `Line 1


Line 2`,
			},
			want: `Line 1

Line 2`,
		},
		{
			name: "Paragraph starts with blank lines",
			args: args{
				s: `
Line 1


Line 2`,
			},
			want: `Line 1

Line 2`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimBlankLines(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestTrimAllLineEndingChars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "HTML String",
			args: args{
				s: `<p>Paragraph 1</p>
<p>Paragraph 2</p>
<p>Paragraph 3</p>`,
			},
			want: "<p>Paragraph 1</p><p>Paragraph 2</p><p>Paragraph 3</p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimAllLineEndingChars(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
