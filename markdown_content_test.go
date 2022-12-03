package md2html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mdcInputArgs struct {
	markdown         string
	excerptSeparator string
}

type mdcOutputArgs struct {
	title       string
	frontmatter map[string]interface{}
	markdown    string
	html        string
}

type mdcTestArgs struct {
	name   string
	input  mdcInputArgs
	output mdcOutputArgs
}

var mdcTests = []mdcTestArgs{
	{
		name: "markdownContent with full options",
		input: mdcInputArgs{
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
			excerptSeparator: "<!-- more -->",
		},
		output: mdcOutputArgs{
			title: "Post title",
			frontmatter: map[string]interface{}{
				"layout":  "default",
				"slug":    "post-slug",
				"author":  "Fake Name",
				"pubDate": "2022-01-01T00:00:00+07:00",
				"keywords": []interface{}{
					"key word 01",
					"key word 02",
				},
				"customKey": "custom value",
				"customArray": []interface{}{
					"element 01",
					"element 02",
				},
				"excerpt": "Post excerpt",
			},
			markdown: `# An h1 header

Paragraphs are separated by a blank line.
`,
			html: `  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>`,
		},
	},
	{
		name: "markdownContent with full options and frontmatter has comments",
		input: mdcInputArgs{
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
			excerptSeparator: "<!-- more -->",
		},
		output: mdcOutputArgs{
			title: "Post title",
			frontmatter: map[string]interface{}{
				"layout":  "default",
				"slug":    "post-slug",
				"author":  "Fake Name",
				"pubDate": "2022-01-01T00:00:00+07:00",
				"keywords": []interface{}{
					"key word 01",
					"key word 02",
				},
				"customKey": "custom value",
				"customArray": []interface{}{
					"element 01",
					"element 02",
				},
				"excerpt": "Post excerpt",
			},
			markdown: `# An h1 header

Paragraphs are separated by a blank line.
`,
			html: `  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>`,
		},
	},
	{
		name: "markdownContent without frontmatter",
		input: mdcInputArgs{
			markdown: `# An h1 header

Paragraphs are separated by a blank line.
`,
			excerptSeparator: "",
		},
		output: mdcOutputArgs{
			title:       "An h1 header",
			frontmatter: map[string]interface{}{},
			markdown: `# An h1 header

Paragraphs are separated by a blank line.
`,
			html: `  <h1 id="an-h1-header">An h1 header</h1>
  <p>Paragraphs are separated by a blank line.</p>`,
		},
	},
}

func Test_markdownContent_Frontmatter(t *testing.T) {
	for _, tt := range mdcTests {
		t.Run(tt.name, func(t *testing.T) {
			mdc, err := newMarkdownContent([]byte(tt.input.markdown), tt.input.excerptSeparator)

			require.NoError(t, err)
			for k, v := range tt.output.frontmatter {
				require.Equal(t, v, mdc.field(k))
			}
		})
	}
}

func Test_markdownContent_Title(t *testing.T) {
	for _, tt := range mdcTests {
		t.Run(tt.name, func(t *testing.T) {
			mdc, err := newMarkdownContent([]byte(tt.input.markdown), tt.input.excerptSeparator)

			require.NoError(t, err)
			require.Equal(t, tt.output.title, mdc.title())
		})
	}
}

func Test_markdownContent_Markdown(t *testing.T) {
	for _, tt := range mdcTests {
		t.Run(tt.name, func(t *testing.T) {
			mdc, err := newMarkdownContent([]byte(tt.input.markdown), tt.input.excerptSeparator)

			require.NoError(t, err)
			require.Equal(t, tt.output.markdown, mdc.markdown())
		})
	}
}

func Test_markdownContent_HTML(t *testing.T) {
	for _, tt := range mdcTests {
		t.Run(tt.name, func(t *testing.T) {
			mdc, err := newMarkdownContent([]byte(tt.input.markdown), tt.input.excerptSeparator)

			require.NoError(t, err)
			require.Equal(t, tt.output.html, mdc.html())
		})
	}
}
