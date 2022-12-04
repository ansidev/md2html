package main

import (
	"fmt"
	"os"

	"github.com/ansidev/md2html"
)

func main() {
	example1()
	example2()
}

func example1() {
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

func example2() {
	s := `---
title: Title 1
slug: title-1
author: John Doe
---

Post excerpt

<!-- more -->

# Heading 1

## Heading 2

Paragraph 1

Paragraph 2`

	tmpl, _ := os.ReadFile("template.html")
	html, err := md2html.Convert([]byte(s), md2html.Options{
		ExcerptSeparator: "<!-- more -->",
		Style:            "light",
		Minify:           false,
		HTMLTemplate:     string(tmpl),
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(html)
}
