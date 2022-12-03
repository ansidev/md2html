package main

import (
	"fmt"

	"github.com/ansidev/md2html"
)

func main() {
	s := `---
title: Title 1
slug: title-1
---

Post excerpt

<!-- more -->

# Heading 1

## Heading 2

Paragraph 1

Paragraph 2`

	html, err := md2html.Convert([]byte(s), md2html.Options{
		ExcerptSeparator: "<!-- more -->",
		Style:            "light",
		Minify:           true,
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(html)
}
