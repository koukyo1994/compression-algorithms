package huffman

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Next   *Node
	Freq   uint64
	Char   rune
}

type Code struct {
	Char  rune
	Value int
	Bit   rune
	Next  *Code
}
