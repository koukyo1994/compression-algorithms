package huffman

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Freq   uint64
	Char   rune
}

type Code struct {
	Char  rune
	Value int
	Bit   rune
	Next  *Code
}

type NodesHeap []Node

func (h NodesHeap) Len() int {
	return len(h)
}

func (h NodesHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h NodesHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodesHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
