package huffman

func BuildCodeBook(tree NodesHeap) map[rune]int {
	codeBook := make(map[rune]int)
	for _, node := range tree {
		// 葉ノードのみ探索を行う
		if node.Left == nil && node.Right == nil {
			_, code, _ := buildCode(node, 0, 0)
			codeBook[node.Char] = code
		}
	}
	return codeBook
}

func buildCode(node *Node, code int, bit int) (*Node, int, int) {
	if node.Which {
		code = code + (1 << bit)
	}
	bit++
	if node.Parent == nil {
		return nil, code, bit
	} else {
		return buildCode(node.Parent, code, bit)
	}
}
