func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {

var dfs func(int,int)bool
adj_list:=make([][]int,numCourses)
for i:=0;i<numCourses;i++{
adj_list[i]=[]int{}
}

// 0-->1-->
for _,pre:=range prerequisites{
	crs,prereq:=pre[0],pre[1]
	adj_list[crs]=append(adj_list[crs],prereq)
}


dfs = func(node int,target int)bool{
	if node==target{
		return true
	}
	for _,nei:=range adj_list[node]{
		if dfs(nei,target){
return true
		}
           
		
	}
	return false
}

res:=make([]bool,0)
for _,q:=range queries{
ans:=dfs(q[0],q[1])
res=append(res,ans)
}

return res

}
