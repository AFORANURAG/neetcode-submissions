func rob(nums []int) int {
    // at any point, i could pick only one 
  n:=len(nums)
  if n==1{
    return nums[0]
  }

  if n==2{
    return max(nums[0],nums[1])
  }
  
  m1:=max(nums[0],nums[1])
  prev:=nums[0]
  curr:=m1

  for i:=2;i<n;i++{
    temp:=curr
    curr=max(curr,nums[i]+prev)
    prev = temp
  }



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
  return curr
}



func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}