func tribonacci(n int) int {


	if n==0{
		return 0
	}
	 if n==1 || n==2 {
		return 1
	 }
	

// memo:=make(map[int]int)
dp:=make([]int,n+1)
dp[0]=0
dp[1]=1
dp[2]=1

for i:=3;i<=n;i++{
	dp[i]=dp[i-1]+dp[i-2]+dp[i-3]
}
return dp[n]
	// var dfs func(int)int
	// dfs = func(k int)int{

    //  if val,ok:=memo[k];ok{
	// 	return val
	//  }
    //  if k==0{
	// 	return 0
	//  }
	//  if k==1 || k==2 {
	// 	return 1
	//  }
	//  v:= dfs(k-1)+dfs(k-2)+dfs(k-3)
	//  memo[k]=v
	//  return v
	// }

	// return dfs(n)
}

// Tn = tn-3+tn-2+tn-1