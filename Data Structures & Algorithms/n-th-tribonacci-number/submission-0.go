func tribonacci(n int) int {
	

memo:=make(map[int]int)
	var dfs func(int)int
	dfs = func(k int)int{

     if val,ok:=memo[k];ok{
		return val
	 }
     if k==0{
		return 0
	 }
	 if k==1 || k==2 {
		return 1
	 }
	 v:= dfs(k-1)+dfs(k-2)+dfs(k-3)
	 memo[k]=v
	 return v
	}

	return dfs(n)
}

// Tn = tn-3+tn-2+tn-1