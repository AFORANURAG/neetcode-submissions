type DisjointSetUnion struct {
	Parent []int
	Rank   []int
	Count int
}

// initialize the DSU with n elements
func  NewDisjointSetUnion(n int) *DisjointSetUnion {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSetUnion{
		Parent: parent,
		Rank:   rank,
		Count:n,
	}

}

func (dsu *DisjointSetUnion) Find(x int) int {
	// either the parent is the same as the element or some other element will be parent of the element
	// base case is when the parent is the same as the element
	// if not, we need to find the parent of the parent
	if dsu.Parent[x] != x {
		dsu.Parent[x] = dsu.Find(dsu.Parent[x])
	} // path compression
	return dsu.Parent[x]
}

// union by rank

func (dsu *DisjointSetUnion) Union(x, y int) bool {
    parentX := dsu.Find(x)
    parentY := dsu.Find(y)

    if parentX == parentY {
        return false
    }

    if dsu.Rank[parentX] < dsu.Rank[parentY] {
        parentX, parentY = parentY, parentX
    }

    // parentX has the larger component
    dsu.Parent[parentY] = parentX
    dsu.Rank[parentX] += dsu.Rank[parentY]
	dsu.Count-=1

    return true
}



func canTraverseAllPairs(nums []int) bool {
	// lets build every pair
	// The problem statement can be boiled down to  are all components connected ?
	l :=len(nums) 
	uf := NewDisjointSetUnion(l)
	factor_idx:=make(map[int]int)

	for i,n:=range nums{
		f := 2

		for f*f<=n{
			if n%f==0{
				if val,ok:=factor_idx[f];ok{
					uf.Union(i,val)
				}else{
					factor_idx[f] = i
				}

			for n%f==0{
				n = n/f
			}
			}
			
		f++
		}

		if n>1{
			if val,ok:=factor_idx[n];ok{
				uf.Union(i,val)
			}else{
				factor_idx[n] = i
			}
		}

	}
	fmt.Println("count is",uf.Count)
	return uf.Count==1
}
