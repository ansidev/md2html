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

func LineEnding() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	// return "\r\n"
	return "\n"
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

func AppendIndentation(s string, indentString string) string {
	sc := bufio.NewScanner(strings.NewReader(s))
	var buf strings.Builder
	isFirstLine := true
	for sc.Scan() {
		if isFirstLine {
			buf.WriteString(indentString)
			buf.WriteString(sc.Text())
			isFirstLine = false
		} else {
			buf.WriteString(LineEnding())
			buf.WriteString(indentString)
			buf.WriteString(sc.Text())
		}
	}

	return buf.String()
}

func TrimBlankLines(s string) string {
	lineEndingRegex := lineEndingUnixRegex
	lineEnding := LineEnding()
	if runtime.GOOS == "windows" {
		lineEndingRegex = lineEndingWindowsRegex
	}

	s1 := lineEndingRegex.ReplaceAllString(s, lineEnding)

	return strings.TrimPrefix(s1, lineEnding)
}

func TrimAllLineEndingChars(s string) string {
	return strings.ReplaceAll(s, LineEnding(), "")
}
