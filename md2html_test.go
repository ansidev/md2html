package md2html

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ansidev/md2html/utils"
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

const stylePathPattern string = "styles/github-%s.css"

var convertTests = []convertTestArgs{
	{
		name: "input markdown with full options",
		input: convertInputArgs{
			markdown: `---{{EOL}}layout: default{{EOL}}title: Post title{{EOL}}slug: post-slug{{EOL}}author: Fake Name{{EOL}}pubDate: "2022-01-01T00:00:00+07:00"{{EOL}}keywords:{{EOL}}- key word 01{{EOL}}- key word 02{{EOL}}customKey: custom value{{EOL}}customArray:{{EOL}}- element 01{{EOL}}- element 02{{EOL}}---{{EOL}}{{EOL}}Post excerpt{{EOL}}{{EOL}}<!-- more -->{{EOL}}{{EOL}}# An h1 header{{EOL}}{{EOL}}Paragraphs are separated by a blank line.{{EOL}}`,
			options: Options{
				ExcerptSeparator: "<!-- more -->",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>{{EOL}}<head>{{EOL}}  <meta charset="utf-8">{{EOL}}  <meta name="viewport" content="width=device-width, initial-scale=1">{{EOL}}  <title>Post title</title>{{EOL}}  <style>{{.Style}}</style>{{EOL}}  <style>body { padding : 15px; }</style>{{EOL}}</head>{{EOL}}{{EOL}}<body class="markdown-body">{{EOL}}<h1 id="an-h1-header">An h1 header</h1>{{EOL}}<p>Paragraphs are separated by a blank line.</p>{{EOL}}</body>{{EOL}}</html>{{EOL}}`,
		},
	},
	{
		name: "input markdown with full options and frontmatter has comments",
		input: convertInputArgs{
			markdown: `---{{EOL}}# Layout name{{EOL}}layout: default{{EOL}}title: Post title # post title{{EOL}}slug: post-slug{{EOL}}author: Fake Name{{EOL}}pubDate: "2022-01-01T00:00:00+07:00"{{EOL}}keywords:{{EOL}}- key word 01{{EOL}}- key word 02{{EOL}}customKey: custom value{{EOL}}customArray:{{EOL}}- element 01{{EOL}}- element 02{{EOL}}---{{EOL}}{{EOL}}Post excerpt{{EOL}}{{EOL}}<!-- more -->{{EOL}}{{EOL}}# An h1 header{{EOL}}{{EOL}}Paragraphs are separated by a blank line.{{EOL}}`,
			options: Options{
				ExcerptSeparator: "<!-- more -->",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>{{EOL}}<head>{{EOL}}  <meta charset="utf-8">{{EOL}}  <meta name="viewport" content="width=device-width, initial-scale=1">{{EOL}}  <title>Post title</title>{{EOL}}  <style>{{.Style}}</style>{{EOL}}  <style>body { padding : 15px; }</style>{{EOL}}</head>{{EOL}}{{EOL}}<body class="markdown-body">{{EOL}}<h1 id="an-h1-header">An h1 header</h1>{{EOL}}<p>Paragraphs are separated by a blank line.</p>{{EOL}}</body>{{EOL}}</html>{{EOL}}`,
		},
	},
	{
		name: "input markdown without frontmatter",
		input: convertInputArgs{
			markdown: `# An h1 header{{EOL}}{{EOL}}Paragraphs are separated by a blank line.{{EOL}}`,
			options: Options{
				ExcerptSeparator: "",
				Style:            "light",
				Minify:           false,
			},
		},
		output: convertOutputArgs{
			htmlWithoutStyle: `<html>{{EOL}}<head>{{EOL}}  <meta charset="utf-8">{{EOL}}  <meta name="viewport" content="width=device-width, initial-scale=1">{{EOL}}  <title>An h1 header</title>{{EOL}}  <style>{{.Style}}</style>{{EOL}}  <style>body { padding : 15px; }</style>{{EOL}}</head>{{EOL}}{{EOL}}<body class="markdown-body">{{EOL}}<h1 id="an-h1-header">An h1 header</h1>{{EOL}}<p>Paragraphs are separated by a blank line.</p>{{EOL}}</body>{{EOL}}</html>{{EOL}}`,
		},
	},
}

func Test_Convert(t *testing.T) {
	for _, tt := range convertTests {
		t.Run(tt.name, func(t *testing.T) {
			markdown := utils.OSBasedStr(tt.input.markdown)
			html, err := Convert([]byte(markdown), tt.input.options)
			require.NoError(t, err)

			b, err1 := os.ReadFile(fmt.Sprintf(stylePathPattern, tt.input.options.Style))
			require.NoError(t, err1)
			require.True(t, len(b) > 0, "Style must have content")

			htmlWithoutStyle := utils.OSBasedStr(tt.output.htmlWithoutStyle)
			expectedHtml := strings.Replace(htmlWithoutStyle, "{{.Style}}", string(b), 1)

			require.Equal(t, expectedHtml, html)
		})
	}
}
