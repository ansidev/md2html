package md2html

const defaultHTMLTemplate = `<html>{{EOL}}<head>{{EOL}}  <meta charset="utf-8">{{EOL}}  <meta name="viewport" content="width=device-width, initial-scale=1">{{EOL}}  <title>{{.Title}}</title>{{EOL}}  <style>{{.Style}}</style>{{EOL}}  <style>body { padding : 15px; }</style>{{EOL}}</head>{{EOL}}{{EOL}}<body class="markdown-body">{{EOL}}{{.Body}}{{EOL}}</body>{{EOL}}</html>{{EOL}}`

type htmlTemplateContext struct {
	Title string
	Body  string
	Style string
}
