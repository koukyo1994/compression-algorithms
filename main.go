package main

import (
	"flag"

	"github.com/koukyo1994/compression-algorithms/runlength"
)

func main() {
	var (
		algorithm = flag.String("algorithm", "run_length", "Algorithm to use for encoding")
		input     = flag.String("input", "", "Input file")
		output    = flag.String("output", "", "Output file")
	)
	flag.Parse()
	if *algorithm == "run_length" {
		if err := runlength.RunLengthEncode(*input, *output); err != nil {
			panic(err)
		}
	} else {
		panic("Unknown algorithm")
	}
}
