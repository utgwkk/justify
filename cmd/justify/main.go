package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/maxmclau/gput"
	"github.com/utgwkk/justify"
)

func main() {
	isAsciiArt := flag.Bool("aa", false, "Enable ascii art mode")
	flag.Parse()
	cols := gput.Cols()

	if *isAsciiArt {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		input := string(buf)
		fmt.Println(justify.CenterAsciiArt(cols, input))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(justify.Center(cols, scanner.Text()))
		}
	}
}
