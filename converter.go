package md2html

import (
	"bytes"
	"text/template"

	"github.com/ansidev/md2html/utils"
)

func compileToHTML(templateCtx htmlTemplateContext, options Options) ([]byte, error) {
	htmlTemplate := options.HTMLTemplate
	if len(htmlTemplate) == 0 {
		htmlTemplate = utils.OSBasedStr(defaultHTMLTemplate)
	}

	tmpl, err := template.New("markdown").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	var compiledHTML bytes.Buffer
	err = tmpl.Execute(&compiledHTML, templateCtx)
	if err != nil {
		return nil, err
	}

	b := compiledHTML.Bytes()
	if !options.Minify {
		return b, nil
	}

	minifiedHTMLBytes, err := minify(b)
	if err != nil {
		return nil, err
	}

	return minifiedHTMLBytes, nil
}
