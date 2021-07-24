package huffman

import (
	"container/heap"
)

func BuildHuffmanTree(nodes NodesHeap) NodesHeap {
	var tree NodesHeap
	for {
		n1 := heap.Pop(&nodes).(*Node)
		n2 := heap.Pop(&nodes).(*Node)

		// Charはいらないが、nilにできないので空文字を入れておく
		// Whichはいらないが、nilにできないのでfalseにする
		parent := Node{Parent: nil, Left: n1, Right: n2, Char: ' ', Freq: n1.Freq + n2.Freq, Which: false}

		n1.Parent = &parent
		n2.Parent = &parent

		n1.Which = false
		n2.Which = true

		heap.Push(&nodes, &parent)
		heap.Push(&tree, n1)
		heap.Push(&tree, n2)

		if nodes.Len() == 1 {
			break
		}
	}

	// 最後のノードはrootノード
	root := heap.Pop(&nodes).(*Node)
	heap.Push(&tree, root)
	return tree
}
