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
