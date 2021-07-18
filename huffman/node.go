package huffman

import (
	"container/heap"
)

func NewNode(char rune, freq uint64) Node {
	return Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Char:   char,
		Freq:   freq,
		Which:  false,
	}
}

func PrepareLeafNodes(charCounts map[rune]int) NodesHeap {
	var leafNodes NodesHeap
	for char, count := range charCounts {
		node := NewNode(char, uint64(count))
		heap.Push(&leafNodes, &node)
	}
	return leafNodes
}
