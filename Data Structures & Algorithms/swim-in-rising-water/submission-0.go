type HeapItem struct {
	Cell []int
	Cost int
}

type PriorityQueue []*HeapItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]

	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*HeapItem)
	*pq = append(*pq, item)
}

func swimInWater(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	pq := &PriorityQueue{}
	heap.Init(pq)

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	heap.Push(pq, &HeapItem{
		Cell: []int{0, 0},
		Cost: grid[0][0],
	})

	for pq.Len() > 0 {

		curr := heap.Pop(pq).(*HeapItem)

		row := curr.Cell[0]
		col := curr.Cell[1]

		if visited[row][col] {
			continue
		}

		visited[row][col] = true

		// reached destination
		if row == n-1 && col == m-1 {
			return curr.Cost
		}

		for _, dir := range directions {

			dr, dc := dir[0], dir[1]

			newR := row + dr
			newC := col + dc

			if newR < 0 || newR >= n || newC < 0 || newC >= m {
				continue
			}

			if visited[newR][newC] {
				continue
			}

			cell := grid[newR][newC]

			heap.Push(pq, &HeapItem{
				Cell: []int{newR, newC},
				Cost: max(curr.Cost, cell),
			})
		}
	}

	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}