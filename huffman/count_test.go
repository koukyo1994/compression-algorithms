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
		{"境界値/空ファイル", "../testdata/empty.txt", map[rune]int{}},
		{"境界値/マルチバイト文字入り", "../testdata/multi_byte.txt", map[rune]int{' ': 6, 'T': 1, 'h': 2, 'i': 4, 's': 4, 'a': 1, 'e': 6, 'n': 2, 't': 6, 'c': 1, 'w': 1, 'm': 1, 'u': 1, 'l': 2, '-': 1, 'b': 1, 'y': 1, 'r': 1, '💀': 1, '.': 1, '\n': 2, 'こ': 1, 'の': 1, 'フ': 1, 'ァ': 1, 'イ': 2, 'ル': 2, 'は': 1, 'マ': 1, 'チ': 1, 'バ': 1, 'ト': 1, '文': 1, '字': 1, 'を': 1, '含': 1, 'み': 1, 'ま': 1, 'す': 1, '😇': 1, '。': 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.filePath)
			assert.Nil(t, err)
			defer file.Close()
			actual := huffman.CountChars(file)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
