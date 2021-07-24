package main

import (
	"fmt"
	"os"

	"github.com/icza/bitio"
)

func main() {
	file, err := os.Create("bitio.dat")
	if err != nil {
		panic(err)
	}

	writer := bitio.NewWriter(file)
	values := []uint64{0x08, 0x07, 0x05, 0x15}
	bits := []uint8{4, 3, 3, 6}
	for i, b := range bits {
		value := values[i]
		if err := writer.WriteBits(value, b); err != nil {
			panic(err)
		}
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}

	file.Close()

	file, err = os.Open("bitio.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bitio.NewReader(file)
	for _, b := range bits {
		value, err := reader.ReadBits(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", value)
	}
}
