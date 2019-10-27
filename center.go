package justify

import (
	"github.com/mattn/go-runewidth"
	"strings"
)

func centerLine(cols, lineWidth int, line string) string {
	if lineWidth >= cols {
		return line
	}

	if lineWidth == 0 {
		return line
	}

	leftPad := (cols - lineWidth) / 2
	padded := strings.Repeat(" ", leftPad) + line

	return padded
}

// Center centerizes normal text in the terminal of `cols` columns.
func Center(cols int, text string) string {
	lines := strings.Split(text, "\n")
	centerizedLines := make([]string, len(lines))

	for i, line := range lines {
		centerizedLines[i] = centerLine(cols, runewidth.StringWidth(line), line)
	}

	return strings.Join(centerizedLines, "\n")
}

// CenterASCIIArt centerizes an ASCII Art in the terminal of `cols` columns.
func CenterASCIIArt(cols int, text string) string {
	lines := strings.Split(text, "\n")
	centerizedLines := make([]string, len(lines))
	maxLength := 0

	for _, line := range lines {
		runeLength := runewidth.StringWidth(line)
		if runeLength > maxLength {
			maxLength = runeLength
		}
	}

	for i, line := range lines {
		centerizedLines[i] = centerLine(cols, maxLength, line)
	}

	return strings.Join(centerizedLines, "\n")
}
