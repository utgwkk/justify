package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/utgwkk/justify"
)

func main() {
	isASCIIArt := flag.Bool("aa", false, "Enable ascii art mode")
	flag.Parse()
	cols := getCols(os.Stdout.Fd(), 80)

	if *isASCIIArt {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		input := string(buf)
		fmt.Println(justify.CenterASCIIArt(cols, input))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(justify.Center(cols, scanner.Text()))
		}
	}
}
