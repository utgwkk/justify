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

func Center (cols int, text string) string {
	lines := strings.Split(text, "\n")
	centerizedLines := make([]string, len(lines))

	for i, line := range lines {
		centerizedLines[i] = centerLine(cols, line)
	}

	return strings.Join(centerizedLines, "\n")
}
