func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    adj := make([][]int, numCourses)
    for i := range adj {
        adj[i] = []int{}
    }
    for _, pre := range prerequisites {
        adj[pre[0]] = append(adj[pre[0]], pre[1])
    }

    var dfs func(node, target int) bool
    dfs = func(node, target int) bool {
        if node == target {
            return true
        }
        for _, nei := range adj[node] {
            if dfs(nei, target) {
                return true
            }
        }
        return false
    }

    res := make([]bool, len(queries))
    for i, q := range queries {
        res[i] = dfs(q[0], q[1])
    }
    return res
}