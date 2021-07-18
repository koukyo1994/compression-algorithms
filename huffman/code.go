package huffman

func BuildCodeBook(tree NodesHeap) map[rune]string {
	codeBook := make(map[rune]string)
	for _, node := range tree {
		// 葉ノードのみ探索を行う
		if node.Left == nil && node.Right == nil {
			_, code := buildCode(node, "")
			codeBook[node.Char] = code
		}
	}
	return codeBook
}

func buildCode(node *Node, code string) (*Node, string) {
	if !node.Which {
		code = "0" + code
	} else {
		code = "1" + code
	}
	if node.Parent == nil {
		return nil, code
	} else {
		return buildCode(node.Parent, code)
	}
}
