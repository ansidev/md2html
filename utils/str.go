package utils

import (
	"bufio"
	"regexp"
	"runtime"
	"strings"
)

var headingRegex = regexp.MustCompile(`([#]+\s)`)
var lineEndingWindowsRegex = regexp.MustCompile(`(?m)^[\r\n]+`)
var lineEndingUnixRegex = regexp.MustCompile(`(?m)^[\n]+`)

func EOL() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func OSBasedStr(s string) string {
	return strings.ReplaceAll(s, "{{EOL}}", EOL())
}

func GetFirstLine(s string) string {
	sc := bufio.NewScanner(strings.NewReader(s))
	line := s
	for sc.Scan() {
		line = sc.Text()
		break
	}

	return headingRegex.ReplaceAllString(line, "")
}

func TrimBlankLines(s string) string {
	lineEndingRegex := lineEndingUnixRegex
	lineEnding := EOL()
	if runtime.GOOS == "windows" {
		lineEndingRegex = lineEndingWindowsRegex
	}

	s1 := lineEndingRegex.ReplaceAllString(s, lineEnding)

	return strings.TrimPrefix(s1, lineEnding)
}

func TrimAllLineEndingChars(s string) string {
	return strings.ReplaceAll(s, EOL(), "")
}
