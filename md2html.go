package md2html

import (
	"fmt"
	"log"
)

// Options for converting from markdown to HTML
type Options struct {
	HTMLTemplate     string
	ExcerptSeparator string
	Style            string
	Minify           bool
}

// Convert converts markdown to HTML string
func Convert(input []byte, options Options) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	markdownContent, err := newMarkdownContent(input, options.ExcerptSeparator)
	if err != nil {
		return "", err
	}

	style := options.Style
	markdownStyle := supportedStyles["light"]
	if ms, ok := supportedStyles[style]; len(ms) > 0 && ok {
		markdownStyle = ms
	}

	title := markdownContent.title()
	if len(title) == 0 {
		return "", fmt.Errorf("title cannot be an empty string")
	}

	htmlContext := htmlTemplateContext{
		Title:       title,
		Body:        markdownContent.html(),
		Frontmatter: markdownContent.frontmatter(),
		Style:       markdownStyle,
	}

	return render(htmlContext, options)
}

func render(templateCtx htmlTemplateContext, options Options) (string, error) {
	b, err := compileToHTML(templateCtx, options)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(b), nil
}
