package huffman

import (
	"bufio"
	"os"
)

func CountChars(f *os.File) map[rune]int {
	reader := bufio.NewReader(f)
	counter := make(map[rune]int)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		counter[r]++
	}
	return counter
}
