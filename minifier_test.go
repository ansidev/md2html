package md2html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type minifyInputArgs struct {
	s string
}

type minifyOutputArgs struct {
	s   string
	err error
}

type minifyTestArgs struct {
	name   string
	input  minifyInputArgs
	output minifyOutputArgs
}

var testHTMLString = `<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Page title</title>
  <style>
    .markdown-body {
      color: #000000;
    }
  </style>
  <style>body { padding : 15px; }</style>
</head>

<body class="markdown-body">
  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>
</body>
</html>
`

var testMinifiedHTMLString = `<html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Page title</title><style>.markdown-body{color:#000}</style><style>body{padding:15px}</style></head><body class="markdown-body"><h1 id="an-h1-header">An h1 header</h1><p>Paragraphs are separated by a blank line.</p></body></html>`

var minifyTests = []minifyTestArgs{
	{
		name: "Minify",
		input: minifyInputArgs{
			s: testHTMLString,
		},
		output: minifyOutputArgs{
			s:   testMinifiedHTMLString,
			err: nil,
		},
	},
}

func Test_Minify(t *testing.T) {
	for _, tt := range minifyTests {
		t.Run(tt.name, func(t *testing.T) {
			b := []byte(tt.input.s)
			ob, err := minify(b)

			require.Equal(t, tt.output.err, err)
			require.Equal(t, tt.output.s, string(ob))
		})
	}
}
