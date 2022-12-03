package md2html

import (
	"bytes"

	"github.com/tdewolff/minify/v2/html"
	minifier "github.com/tdewolff/minify/v2/minify"
)

func minify(b []byte) ([]byte, error) {
	m := minifier.Default
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepQuotes:       true,
		KeepEndTags:      true,
	})
	var buf bytes.Buffer
	r := bytes.NewReader(b)
	err := m.Minify("text/html", &buf, r)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
