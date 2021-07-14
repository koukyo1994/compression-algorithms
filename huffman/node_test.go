package huffman_test

import (
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
