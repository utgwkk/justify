// +build windows

package main

import (
	"github.com/Azure/go-ansiterm/winterm"
)

func getCols(fd uintptr, defaultCols int) int {
	info, err := winterm.GetConsoleScreenBufferInfo(fd)
	if err != nil {
		// fall back
		return defaultCols
	}
	return int(info.Window.Right - info.Window.Left + 1)
}
