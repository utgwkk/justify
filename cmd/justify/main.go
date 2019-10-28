package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/utgwkk/justify"
)

func main() {
	isASCIIArt := flag.Bool("aa", false, "Enable ascii art mode")
	flag.Parse()
	cols := getCols(os.Stdout.Fd(), 80)

	if *isASCIIArt {
		var inputLines []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputLines = append(inputLines, scanner.Text())
		}
		input := strings.Join(inputLines, "\n")
		fmt.Println(justify.CenterASCIIArt(cols, input))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(justify.Center(cols, scanner.Text()))
		}
	}
}
