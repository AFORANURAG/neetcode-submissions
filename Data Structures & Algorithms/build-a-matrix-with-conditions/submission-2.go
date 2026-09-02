// lets build critical thinking skills
// we will be feeling so good after doing it

// We would need to write a resusable topological sorting function

func topologicalSort(adjlist map[int][]int) []int {
    recursionPath := make(map[int]bool)
    visited := make(map[int]bool)

    topoSortArray := make([]int, 0)

    var dfs func(int) bool

    dfs = func(node int) bool {
        // Node exists in current recursion stack → cycle
        if recursionPath[node] {
            return true
        }

        // Already completely explored
        if visited[node] {
            return false
        }

        recursionPath[node] = true

        for _, nei := range adjlist[node] {
            if dfs(nei) {
                return true
            }
        }

        // Done exploring this node
        recursionPath[node] = false
        visited[node] = true

        // Postorder
        topoSortArray = append(topoSortArray, node)

        return false
    }

    // Graph may be disconnected
    for node := range adjlist {
        if !visited[node] {
            if dfs(node) {
                return nil // cycle → no topological ordering
            }
        }
    }

    // DFS gives reverse topological order
    for i, j := 0, len(topoSortArray)-1; i < j; i, j = i+1, j-1 {
        topoSortArray[i], topoSortArray[j] =
            topoSortArray[j], topoSortArray[i]
    }

    return topoSortArray
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// rest we will do in evening
	// let's build the intuition first
	// above[i] should be strictly above below[i]
	// let's first apply topological sorting
	// first let's apply on example 1 row conditions
	// [2,1,3] --> row topo
	// now on column conditions, the topo sort will be the below
	// [2,3,1] --> col topo

	// iterate on both and fill the map of coordinates
	// therefore 1:[0,2], 2:[1,0], 3:[2,1], or i think
	// [1,1] --> 1, [0,0] --> 2, [2,1] --> 3

	// say we have [
	// 0,0,0
	// 0,0,0
	// 0,0,0
	// ]
	// let's do this

	adj_list_rows := make(map[int][]int)

	for _, val := range rowConditions {
		u, v := val[0], val[1]
		adj_list_rows[u] = append(adj_list_rows[u], v)
	}

	adj_list_cols := make(map[int][]int)

	for _, val := range colConditions {
		u, v := val[0], val[1]
		adj_list_cols[u] = append(adj_list_cols[u], v)
	}

	toporows := topologicalSort(adj_list_rows)
	topocols := topologicalSort(adj_list_cols)

	if len(toporows) == 0 || len(topocols) == 0 {
		return [][]int{}
	}


	// let's create a matrix now

	sol := make([][]int, k)

	for i := 0; i < k; i++ {
		sol[i] = make([]int, k)
	}

	type Position struct {
		row *int
		col *int
	}

	rowcolmap := make(map[int]Position)

	for i := 1; i <= k; i++ {
		rowcolmap[i] = Position{
			row: nil,
			col: nil,
		}
	}

	for row_idx, row_val := range toporows {
		row := row_idx

		rowcolmap[row_val] = Position{
			row: &row,
			col: nil,
		}
	}

	for col_idx, col_val := range topocols {
		col := col_idx

		p := rowcolmap[col_val]
		p.col = &col
		rowcolmap[col_val] = p
	}

	for val, v := range rowcolmap {

		row_idx, col_idx := v.row, v.col

		if row_idx != nil && col_idx != nil {
			sol[*row_idx][*col_idx] = val
		}

		// if only row
		if row_idx != nil && col_idx == nil {
			row := *row_idx

			// assign a column which is free
			for i := 0; i < k; i++ {
				if sol[row][i] == 0 {
					sol[row][i] = val
					break
				}
			}
		}

		// if only col
		if row_idx == nil && col_idx != nil {
			col := *col_idx

			// assign a row which is free
			for i := 0; i < k; i++ {
				if sol[i][col] == 0 {
					sol[i][col] = val
					break
				}
			}
		}
	}

	// assign values which have neither row nor column
	for val, v := range rowcolmap {

		row_idx, col_idx := v.row, v.col

		if row_idx == nil && col_idx == nil {
			// assign random value
			for i := 0; i < k; i++ {
				found := false

				for j := 0; j < k; j++ {
					if sol[i][j] == 0 {
						sol[i][j] = val
						found = true
						break
					}
				}

				if found {
					break
				}
			}
		}
	}

	return sol
}