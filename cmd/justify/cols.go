// +build linux darwin !windows

package main

import (
	"golang.org/x/sys/unix"
)

func getCols(fd uintptr, defaultCols int) int {
	size, err := unix.IoctlGetWinsize(int(fd), unix.TIOCGWINSZ)
	if err != nil {
		// fall back
		return defaultCols
	}
	return int(size.Col)
}
