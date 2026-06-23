func findRedundantConnection(edges [][]int) []int {
    // build the map
	adj_list:=make(map[int][]int,0)
	visited:=make(map[int]bool)



	
	// first time visited = 1
	// second time visited = 2
	// zero/nil means not visited

	var dfs func(int,int)bool

	dfs = func(node,parent int)bool{
		visited[node]=true
	for _,nei:=range adj_list[node]{
		// we get childs
		if !visited[nei]{
			// not visited
			isCyclic:=dfs(nei,node)
			if isCyclic{
				return true
			}
		}else{
			// one time already visited
			if nei!=parent{
				return true
			}
		}
	}
	return false
	}

	for _,nei:=range edges{
		u,v:=nei[0],nei[1]
		adj_list[u]=append(adj_list[u],v)
		adj_list[v]=append(adj_list[v],u)
	    visited=make(map[int]bool)


		if dfs(u,-1){
			return []int{u,v}
		}
	}


	// We have the map ready
	// Now, collect all the edges where collect exists,
	

return edges[0]

}
