func climbStairs(n int) int {

	var dfs func(int)int
	// count combinations 

	memo:=make(map[int]int)
	dfs = func(sum int)int{
		// include i, go further,remove it and try it for next iteration
		// base case
		
		if val,ok:=memo[sum];ok{
			return val
		}
		if sum>n{
			return 0
		}
		if sum==n{
			return 1
		}
		comb:=	dfs(sum+1)+dfs(sum+2)
		memo[sum]=comb
		return comb
	}

return dfs(0);
}