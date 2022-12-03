package md2html

const defaultHTMLTemplate = `<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>{{.Style}}</style>
  <style>body { padding : 15px; }</style>
</head>

<body class="markdown-body">
{{.Body}}
</body>
</html>
`

type htmlTemplateContext struct {
	Title string
	Body  string
	Style string
}
