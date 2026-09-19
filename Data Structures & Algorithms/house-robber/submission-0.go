func rob(nums []int) int {
    // at any point, i could pick only one 
  n:=len(nums)
  memo:=make(map[int]int)
  var dfs func(int)int
  dfs=func(i int)int{
    if val,ok:=memo[i];ok{
      return val
    }
    if i >=n{
		return 0
	}
  c1:=dfs(i+1)
  c2:=dfs(i+2)
  memo[i]=max(nums[i]+c2,c1)

return memo[i]
  }
  return dfs(0)
}



func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}