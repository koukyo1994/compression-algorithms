package huffman_test

import (
	"container/heap"
	"testing"

	"github.com/koukyo1994/compression-algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestNodesHeap(t *testing.T) {
	tests := []struct {
		name          string
		nodes         []huffman.Node
		expectedFreqs []uint64
	}{
		{
			"正常系",
			[]huffman.Node{
				{Freq: 10, Char: 'a'},
				{Freq: 20, Char: 'b'},
				{Freq: 5, Char: 'c'},
				{Freq: 8, Char: 'd'},
				{Freq: 36, Char: 'e'},
				{Freq: 4, Char: 'f'},
			},
			[]uint64{4, 5, 8, 10, 20, 36},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodesHeap := &huffman.NodesHeap{}
			for _, n := range tt.nodes {
				heap.Push(nodesHeap, n)
			}

			for i := 0; i < len(tt.nodes); i++ {
				v := heap.Pop(nodesHeap).(huffman.Node)
				assert.Equal(t, tt.expectedFreqs[i], v.Freq, "%s: NodesHeap", tt.name)
			}
		})
	}
}
