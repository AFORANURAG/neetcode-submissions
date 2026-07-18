// func minimumEffortPath(heights [][]int) int {
// // effort is abs(hj-hi)
// // so now we can move up,down,left and right
// // minimize the effort
// // Find all the paths
// // and return minimum effort among the path
// // overall tc will be 4^m*n, at each step(m*n), there exists 4 ways to explore
// // so we can explore all and then we can check minimum effort among all the routes
// // We need to return minimum effort
// // [[[0,1],[1,0]]],[[0,2],[1,1],[2,0]],[[1,2],[2,1]],[[2,2]]]
// }

// import (
//     "container/heap"
// )

type Item struct {
    diff, r, c int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].diff < h[j].diff }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minimumEffortPath(heights [][]int) int {
    rows, cols := len(heights), len(heights[0])
    dist := make([][]int, rows)
    // We have distance to continue on stale values

    for i :=range dist{
        dist[i]=make([]int,cols)
        for j:=range dist[i]{
            dist[i][j] = 1 << 31 - 1
        }
    }
    dist[0][0] = 0
    directions:=[][]int{{1,0},{-1,0},{0,1},{0,-1}}
    minHeap:=&MinHeap{Item{0,0,0}}
    heap.Init(minHeap)

    for minHeap.Len() > 0 {
        curr:=heap.Pop(minHeap).(Item)
        diff, r, c := curr.diff, curr.r, curr.c
        if r==rows-1 && c ==cols-1 {
            return diff
        }
        if dist[r][c]<diff{
            continue
        }

        for _,dir:=range directions{
        newR, newC := r+dir[0], c+dir[1]
        if newR < 0 || newR>=rows || newC<0 || newC>=cols {
            continue
        }
        newDiff:=max(diff,abs(heights[r][c]-heights[newR][newC]))
        if newDiff < dist[newR][newC] {
            dist[newR][newC] = newDiff
            heap.Push(minHeap,Item{newDiff,newR,newC})
        }

    }
}
return 0 
}

