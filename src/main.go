package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("mailing", "Mailing provides a simple and robust email solution to newsletter subscriptions and notices. Implements a variety of algorithms that consist in concurrency tasks managements and fake mail detections.")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
}
