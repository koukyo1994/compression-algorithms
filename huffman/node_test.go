package huffman_test

import (
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
		expected []huffman.Node
	}{
		{"正常系/普通のファイル", "../testdata/normal.txt", []huffman.Node{
			{nil, nil, nil, 4, 'a'},
			{nil, nil, nil, 4, 'b'},
			{nil, nil, nil, 1, 'c'},
			{nil, nil, nil, 2, 'n'},
			{nil, nil, nil, 3, 'i'},
			{nil, nil, nil, 1, 'l'},
			{nil, nil, nil, 4, ' '},
			{nil, nil, nil, 1, '\n'},
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
			assert.ElementsMatch(t, tt.expected, nodes, "%s: PrepareLeafNodes", tt.name)
		})
	}
}
