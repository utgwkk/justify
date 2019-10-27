package justify

import (
	"strings"
	"github.com/mattn/go-runewidth"
)

func centerLine (cols int, line string) string {
	halfwidthCols := cols
	lineWidth := runewidth.StringWidth(line)

	if lineWidth >= halfwidthCols {
		return line
	}

	if lineWidth == 0 {
		return line
	}

	leftPad := (halfwidthCols - lineWidth) / 2
	padded := strings.Repeat(" ", leftPad) + line

	return padded
}

func centerLineAA (cols, lineWidth int, line string) string {
	halfwidthCols := cols

	if lineWidth >= halfwidthCols {
		return line
	}

	if lineWidth == 0 {
		return line
	}

	leftPad := (halfwidthCols - lineWidth) / 2
	padded := strings.Repeat(" ", leftPad) + line

	return padded
}

func Center (cols int, text string) string {
	lines := strings.Split(text, "\n")
	centerizedLines := make([]string, len(lines))

	for i, line := range lines {
		centerizedLines[i] = centerLine(cols, line)
	}

	return strings.Join(centerizedLines, "\n")
}

func CenterAsciiArt (cols int, text string) string {
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
		centerizedLines[i] = centerLineAA(cols, maxLength, line)
	}

	return strings.Join(centerizedLines, "\n")
}
