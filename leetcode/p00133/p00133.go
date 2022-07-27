package p00133

type Node struct {
	Val       int
	Neighbors []*Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[*Node]*Node{}
	head := &Node{}
	dfs(visited, node, head)

	return head
}

func dfs(visited map[*Node]*Node, src, dst *Node) {
	dst.Val = src.Val
	visited[src] = dst

	for _, node := range src.Neighbors {
		if reference, exists := visited[node]; exists {
			dst.Neighbors = append(dst.Neighbors, reference)
			continue
		}

		new_node := &Node{}
		dst.Neighbors = append(dst.Neighbors, new_node)
		dfs(visited, node, new_node)
	}
}
