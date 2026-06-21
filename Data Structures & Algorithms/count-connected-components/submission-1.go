func countComponents(n int, edges [][]int) int {
    // how do you calculate number of disconnected components or how many components are there 
    

	// 1:[2,3]
	// 2:[4,6]
	// 3:[9,10]
	// 11:[15,19]
	
	// run dfs starting at 0
	// if len(visited)==n, output 1
	// else iterate on map,and skip all those which have been visited
	// now, you got an edge which has not been visited earlier, now do it again and so on.

	// in actuality, we are iterating two time on each element

	// given: undirected graph
	adj := make([][]int, n)
	visited := make([]bool, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(int)
	dfs = func(node int){
	  visited[node] = true
	  for _,nei:=range adj[node]{
		if !visited[nei]{
			dfs(nei)
		}
	  }	
	}
	var bfs func(int)
	bfs = func(node int){
		q:=make([]int,0)
		visited[node]=true
		q=append(q,node)
		for len(q) > 0{
			curr:=q[0]
			q = q[1:]
			for _,nei:=range adj[curr]{
				if !visited[nei]{
					visited[nei]=true
					q = append(q,nei)
				}
			}
		} 
	}
	res:= 0
	for node := 0; node < n; node++ {
		if !visited[node] {
			bfs(node)
			res++
		}
	}
	return res
}
