func validTree(n int, edges [][]int) bool {
    //  0
  //   1
  //  2
  // 3
  // 1:[2,3,4]
  // 2:[4]

  adj_list:=make(map[int][]int)
  for _,edge:=range edges{
	src,dest:=edge[0],edge[1]
	adj_list[src]=append(adj_list[src],dest)
	adj_list[dest]=append(adj_list[dest],src)
  }
  visited:=make(map[int]bool)

  var dfs func(int,int)bool
  dfs = func(node,pnode int)bool{
    if visited[node]{
      return false
    }
	visited[node]=true
	for _,nei:=range adj_list[node]{
   if nei==pnode{
    continue
   }
   if !dfs(nei,node){
    return false
   }

	}
	return true
  }
  
return dfs(0,-1) && len(visited)==n
}
