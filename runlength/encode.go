package runlength

import (
	"fmt"
	"io/ioutil"
)

func Encode(lines string) string {
	var prev rune
	var count int
	result := ""
	for pos, c := range lines {
		if c == prev {
			count++
		} else {
			if pos == 0 {
			} else if count > 1 {
				result += string(prev) + fmt.Sprint(count)
			} else {
				result += string(prev)
			}
			count = 1
			prev = c
		}
	}
	if count > 1 {
		result += string(prev) + fmt.Sprint(count)
	} else {
		result += string(prev)
	}
	return result
}

func RunLengthEncode(inPath string, outPath string) error {
	buffer, err := ioutil.ReadFile(inPath)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outPath, []byte(Encode(string(buffer))), 0644)
}
