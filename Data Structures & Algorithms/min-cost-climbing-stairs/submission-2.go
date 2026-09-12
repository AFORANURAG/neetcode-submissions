func minCostClimbingStairs(cost []int) int {
    memo := make([]int, len(cost))
    for i := 0; i < len(cost); i++ {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= len(cost) {
            return 0
        }
        if memo[i] != -1 {
            return memo[i]
        }
        memo[i] = cost[i] + min(dfs(i+1), dfs(i+2))
        return memo[i]
    }

    return min(dfs(0), dfs(1))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}