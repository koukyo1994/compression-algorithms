package huffman_test

import (
	"os"
	"testing"

	"github.com/koukyo1994/compression-algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected map[rune]int
	}{
		{"正常系/普通のファイル", "../testdata/normal.txt", map[rune]int{'a': 4, 'b': 4, 'c': 1, 'n': 2, 'i': 3, 'l': 1, ' ': 4, '\n': 1}},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.filePath)
		if err != nil {
			t.Errorf("%s: Open: %v", tt.name, err)
			continue
		}
		defer file.Close()

		actual := huffman.CountChars(file)
		assert.Equal(t, tt.expected, actual, "%s: CountChars", tt.name)
	}
}
