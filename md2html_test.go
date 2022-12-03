package md2html

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type convertInputArgs struct {
	markdown string
	options  Options
}

type convertOutputArgs struct {
	htmlWithoutStyle string
}

type convertTestArgs struct {
	name   string
	input  convertInputArgs
	output convertOutputArgs
}

const stylePathPattern string = "styles/github-markdown-css/dist/github-%s.css"

var convertTests = []convertTestArgs{
	{
		name: "markdownContent with full options",
		input: convertInputArgs{
			markdown: `---
layout: default
title: Post title
slug: post-slug
author: Fake Name
pubDate: "2022-01-01T00:00:00+07:00"
keywords:
- key word 01
- key word 02
customKey: custom value
customArray:
- element 01
- element 02
---

Post excerpt

<!-- more -->

# An h1 header

Paragraphs are separated by a blank line.
`,
			options: Options{
				ExcerptSeparator: "<!-- more -->",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Post title</title>
  <style>{{.Style}}</style>
  <style>body { padding : 15px; }</style>
</head>

<body class="markdown-body">
  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>
</body>
</html>
`,
		},
	},
	{
		name: "markdownContent with full options and frontmatter has comments",
		input: convertInputArgs{
			markdown: `---
# Layout name
layout: default
title: Post title # post title
slug: post-slug
author: Fake Name
pubDate: "2022-01-01T00:00:00+07:00"
keywords:
- key word 01
- key word 02
customKey: custom value
customArray:
- element 01
- element 02
---

Post excerpt

<!-- more -->

# An h1 header

Paragraphs are separated by a blank line.
`,
			options: Options{
				ExcerptSeparator: "<!-- more -->",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Post title</title>
  <style>{{.Style}}</style>
  <style>body { padding : 15px; }</style>
</head>

<body class="markdown-body">
  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>
</body>
</html>
`,
		},
	},
	{
		name: "markdownContent without frontmatter",
		input: convertInputArgs{
			markdown: `# An h1 header

Paragraphs are separated by a blank line.
`,
			options: Options{
				ExcerptSeparator: "",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>An h1 header</title>
  <style>{{.Style}}</style>
  <style>body { padding : 15px; }</style>
</head>

<body class="markdown-body">
  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>
</body>
</html>
`,
		},
	},
}

func Test_Convert(t *testing.T) {
	for _, tt := range convertTests {
		t.Run(tt.name, func(t *testing.T) {
			html, err := Convert([]byte(tt.input.markdown), tt.input.options)
			require.NoError(t, err)

			b, err1 := os.ReadFile(fmt.Sprintf(stylePathPattern, tt.input.options.Style))
			require.NoError(t, err1)
			require.True(t, len(b) > 0, "Style must have content")

			expectedHtml := strings.Replace(tt.output.htmlWithoutStyle, "{{.Style}}", string(b), 1)

			require.Equal(t, expectedHtml, html)
		})
	}
}
