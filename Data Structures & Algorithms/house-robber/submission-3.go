func rob(nums []int) int {
    // at any point, i could pick only one 
  n:=len(nums)
  if n==1{
    return nums[0]
  }

  if n==2{
    return max(nums[0],nums[1])
  }
  dp:=make(map[int]int)
  
  m1:=max(nums[0],nums[1])
  dp[0]=nums[0]
  dp[1]=m1

  for i:=2;i<n;i++{
    dp[i]=max(dp[i-1],nums[i]+dp[i-2])
  }

  fmt.Println("dp is",dp)


//   var dfs func(int)int
//   dfs=func(i int)int{
//     if val,ok:=memo[i];ok{
//       return val
//     }
//     if i >=n{
// 		return 0
// 	}
//   c1:=dfs(i+1)
//   c2:=dfs(i+2)
//   memo[i]=max(nums[i]+c2,c1)

// return memo[i]
//   }
  return dp[n-1]
}



func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}