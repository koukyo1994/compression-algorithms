package huffman_test

import (
	"os"
	"testing"

	"github.com/koukyo1994/compression-algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestBuildCodeBook(t *testing.T) {
	tests := []struct {
		name             string
		filePath         string
		expectedNumCodes int
	}{
		{
			name:             "正常系",
			filePath:         "../../testdata/normal.txt",
			expectedNumCodes: 8,
		},
		{
			name:             "正常系2",
			filePath:         "../../testdata/abcd.txt",
			expectedNumCodes: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.filePath)
			assert.Nil(t, err, "%s: Open", tt.name)
			defer file.Close()

			charCounts := huffman.CountChars(file)
			nodes := huffman.PrepareLeafNodes(charCounts)

			tree := huffman.BuildHuffmanTree(nodes)
			codeBook := huffman.BuildCodeBook(tree)

			assert.Equal(t, tt.expectedNumCodes, len(codeBook), "%s: len(codeBook)", tt.name)
		})
	}
}
