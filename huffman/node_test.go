package huffman_test

import (
	"container/heap"
	"os"
	"testing"

	"github.com/koukyo1994/compression-algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		name string
		char rune
		freq uint64
	}{
		{"正常系", 'a', 5},
	}

	for _, tt := range tests {
		node := huffman.NewNode(tt.char, tt.freq)
		assert.IsType(t, huffman.Node{}, node, "%s: NewNode", tt.name)
	}
}

func TestPrepareLeafNodes(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected []*huffman.Node
	}{
		{"正常系/普通のファイル", "../testdata/normal.txt", []*huffman.Node{
			{nil, nil, nil, 1, 'c', false},
			{nil, nil, nil, 1, 'l', false},
			{nil, nil, nil, 1, '\n', false},
			{nil, nil, nil, 2, 'n', false},
			{nil, nil, nil, 3, 'i', false},
			{nil, nil, nil, 4, 'a', false},
			{nil, nil, nil, 4, 'b', false},
			{nil, nil, nil, 4, ' ', false},
		}},
		{"境界値/空ファイル", "../testdata/empty.txt", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.filepath)
			assert.Nil(t, err, "%s: Open", tt.name)
			defer file.Close()
			charCounts := huffman.CountChars(file)
			nodes := huffman.PrepareLeafNodes(charCounts)

			expected := huffman.NodesHeap{}
			for _, node := range tt.expected {
				heap.Push(&expected, node)
			}

			assert.Equal(t, expected.Len(), nodes.Len(), "%s: Len", tt.name)
			lenNodes := nodes.Len()
			for i := 0; i < lenNodes; i++ {
				actualNode := heap.Pop(&nodes).(*huffman.Node)
				expectedNode := heap.Pop(&expected).(*huffman.Node)
				assert.Equal(t, expectedNode.Freq, actualNode.Freq, "%s: Pop", tt.name)
			}
		})
	}
}
