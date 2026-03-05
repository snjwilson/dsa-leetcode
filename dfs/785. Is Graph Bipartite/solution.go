package main

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := range graph {
		if colors[i] == 0 {
			if !dfs(i, graph, 1, colors) {
				return false
			}
		}
	}
	return true
}

func dfs(node int, graph [][]int, color int, colors []int) bool {
	if colors[node] != 0 {
		return colors[node] == color
	}
	colors[node] = color
	for _, curr := range graph[node] {
		if !dfs(curr, graph, -color, colors) {
			return false
		}
	}
	return true
}
