func findOrder(numCourses int, prerequisites [][]int) []int {
    // We can do it using kahn's algo(cycle detection with bfs)
	// We can do it using cycle detection in cyclic graph using dfs

// build the adjacency list
preMap:=make(map[int][]int)
indegree:=make([]int,numCourses)
for i:=0;i<numCourses;i++{
	preMap[i]=[]int{}
}

for _,prereq:=range prerequisites {
	crs,pre:=prereq[0],prereq[1]
	preMap[crs]=append(preMap[crs],pre)
	indegree[pre]++
}

queue:=make([]int,0)
fmt.Println("indegree is",indegree)
for i:=0;i<numCourses;i++{
	if indegree[i]==0{
		queue=append(queue,i)
	}
}
fmt.Println("queue is",queue)
output:=make([]int,0)
finish:=0
for len(queue)>0{
curr:=queue[0]
queue=queue[1:]	
output = append(output, curr)
finish++
for _,nei:=range preMap[curr]{
	// all the childs of curr
	fmt.Println("nei is",nei)
	indegree[nei]--
	if indegree[nei]==0{
		queue=append(queue,nei)
	}
}
}
fmt.Println("output is",output)
fmt.Println("finish is",finish)
if finish!=numCourses{
	return []int{}
}

    for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
        output[i], output[j] = output[j], output[i]
    }
    return output

}

