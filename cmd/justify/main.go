package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maxmclau/gput"
	"github.com/utgwkk/justify"
)

func main() {
	cols := gput.Cols()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(justify.Center(cols, scanner.Text()))
	}
}
