func islandsAndTreasure(grid [][]int) {
    // multi-source bfs

	// put all the 0s in the queue
	// start bfs then
	queue:=make([][2]int,0)
	m,n:=len(grid),len(grid[0])
	for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                queue = append(queue, [2]int{i, j})
            }
        }
    }
	dirs:=[][2]int{{-1,0},{1,0},{0,-1},{0,1}}
	INF:=2147483647

	for len(queue)>0{
		node:=queue[0]
		queue = queue[1:]
		row,col:=node[0],node[1]
		for _,dir:=range dirs{
			dr,dy:=dir[0],dir[1]
			r,c:=row+dr,col+dy
			if c<0 || r<0 || r>=m || c>=n || grid[r][c]!= INF{
				continue
			}
			grid[r][c]=grid[row][col]+1
			queue = append(queue, [2]int{r, c})
		}
	}
}
