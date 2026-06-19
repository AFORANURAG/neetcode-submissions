func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {

	adjList := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		adjList[i] = []int{}
	}

	for _, pre := range prerequisites {
		crs, prereq := pre[0], pre[1]
		adjList[crs] = append(adjList[crs], prereq)
	}

	// memo[(u,v)] = whether v is reachable from u
	memo := make(map[[2]int]bool)
	computed := make(map[[2]int]bool)

	var dfs func(int, int) bool

	dfs = func(node, target int) bool {
		key := [2]int{node, target}

		if computed[key] {
			return memo[key]
		}

		if node == target {
			return true
		}

		computed[key] = true

		for _, nei := range adjList[node] {
			if dfs(nei, target) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	res := make([]bool, len(queries))

	for i, q := range queries {
		res[i] = dfs(q[0], q[1])
	}

	return res
}