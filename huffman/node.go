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
	}
}

func PrepareLeafNodes(charCounts map[rune]int) NodesHeap {
	var leafNodes NodesHeap
	for char, count := range charCounts {
		heap.Push(&leafNodes, NewNode(char, uint64(count)))
	}
	return leafNodes
}
