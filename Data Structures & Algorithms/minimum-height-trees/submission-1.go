func findMinHeightTrees(n int, edges [][]int) []int {
// lets do it recursion that is dfs and we will see the optimized ones later on
 adj := make([][]int, n)
    for i := range adj {
        adj[i] = []int{}
    }
    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], edge[1])
        adj[edge[1]] = append(adj[edge[1]], edge[0])
    }
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		hgt:=0
		for _,nei:=range adj[node]{
			if nei==parent{
				continue
			}
			hgt = max(hgt, 1+dfs(nei, node))
		}
		return hgt
	}
    minHeight:=n
	res:=[]int{}
	for i:=0;i<n;i++{
		// run dfs 
		height:=dfs(i,-1)
		if height==minHeight{
			res = append(res,i)
		}else if height< minHeight{
			res = []int{i}
			minHeight=height
		}
	} 
return res

}
