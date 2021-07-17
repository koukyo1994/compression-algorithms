package huffman_test

import (
	"container/heap"
	"os"
	"testing"

	"github.com/koukyo1994/compression-algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestBuildHuffmanTree(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected []uint64
	}{
		{
			name:     "正常系",
			filePath: "../testdata/normal.txt",
			expected: []uint64{
				1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 7, 8, 12, 20,
			},
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

			assert.Equal(t, len(tt.expected), tree.Len(), "%s: BuildHuffmanTree", tt.name)
			lenTree := tree.Len()
			for i := 0; i < lenTree; i++ {
				node := heap.Pop(&tree).(huffman.Node)
				assert.Equal(t, tt.expected[i], node.Freq, "%s: BuildHuffmanTree", tt.name)
			}
		})
	}
}
