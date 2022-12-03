package md2html

import (
	"bytes"
	"log"
	"strings"

	fm "github.com/adrg/frontmatter"
	"github.com/ansidev/md2html/utils"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

type markdownContent struct {
	frontmatter map[string]interface{}
	content     string
}

func newMarkdownContent(bytes []byte, excerptSeparator string) (*markdownContent, error) {
	frontmatter := make(map[string]interface{})

	contentBytes, err := fm.Parse(strings.NewReader(string(bytes)), &frontmatter)
	if err != nil {
		return nil, err
	}

	hasTitle := true

	if t, ok := frontmatter["title"]; !ok {
		hasTitle = false
	} else {
		t1, ok := t.(string)
		if !ok || len(t1) == 0 {
			hasTitle = false
		}
	}

	content := string(contentBytes)
	if !hasTitle {
		frontmatter["title"] = utils.GetFirstLine(content)
	}

	mdc := &markdownContent{frontmatter, content}
	if len(excerptSeparator) == 0 {
		return mdc, nil
	}

	contentSlice := strings.Split(content, excerptSeparator)

	if len(contentSlice) == 0 {
		return mdc, nil
	}

	frontmatter["excerpt"] = utils.TrimAllLineEndingChars(contentSlice[0])
	content = strings.Join(contentSlice[1:], excerptSeparator)

	return &markdownContent{frontmatter, content}, nil
}

func (c *markdownContent) field(key string) interface{} {
	if key == "title" {
		return c.title()
	}

	if val, ok := c.frontmatter[key]; ok {
		return val
	}

	return nil
}

func (c *markdownContent) title() string {
	if val, ok := c.frontmatter["title"]; ok {
		title, ok1 := val.(string)
		if ok1 {
			return title
		}
	}

	return ""
}

func (c *markdownContent) markdown() string {
	return utils.TrimBlankLines(c.content)
}

func (c *markdownContent) htmlWithIndentation(indentString string) string {
	markdownBytes := []byte(c.content)
	mdconv := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	var buf bytes.Buffer
	if err := mdconv.Convert(markdownBytes, &buf); err != nil {
		log.Fatal(err)
		return ""
	}

	s := strings.TrimSuffix(buf.String(), utils.LineEnding())

	return utils.AppendIndentation(s, indentString)
}

func (c *markdownContent) html() string {
	return c.htmlWithIndentation("  ")
}
