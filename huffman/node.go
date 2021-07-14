package huffman

func NewNode(char rune, freq uint64) Node {
	return Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Char:   char,
		Freq:   freq,
	}
}

func PrepareLeafNodes(charCounts map[rune]int) []Node {
	var leafNodes []Node
	for char, count := range charCounts {
		leafNodes = append(leafNodes, NewNode(char, uint64(count)))
	}
	return leafNodes
}
