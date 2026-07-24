import (
	"slices"
	"cmp"
)
// lets implement a disjoint set union data structure


type DisjointSetUnion struct {
	Parent []int
	Rank   []int
}

// initialize the DSU with n elements
func (dsu *DisjointSetUnion) NewDisjointSetUnion(n int) *DisjointSetUnion {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSetUnion{
		Parent: parent,
		Rank:   rank,
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
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)
	if rootX == rootY {
		return false
	}
	// let me write the algo

	if dsu.Rank[rootX] > dsu.Rank[rootY] {
		dsu.Parent[rootY] = rootX
	} else if dsu.Rank[rootX] < dsu.Rank[rootY] {
		dsu.Parent[rootX] = rootY
	} else {
		// if the ranks are the same, we can choose any one as the parent
		dsu.Parent[rootY] = rootX
		dsu.Rank[rootX]++
	}

	return true
}

func (dsu *DisjointSetUnion) IsConnected(x, y int) bool {
	return dsu.Find(x) == dsu.Find(y)
}


type Edge struct {
	u int
	v int
	d int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)

	// Step 1: Generate all edges
	edges := make([]Edge, 0)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := abs(points[i][0]-points[j][0]) +
				abs(points[i][1]-points[j][1])

			edges = append(edges, Edge{
				u: i,
				v: j,
				d: distance,
			})
		}
	}

	// Step 2: Sort edges by weight
	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.d, b.d)
	})

	// Step 3: Initialize DSU
	dsu := new(DisjointSetUnion).NewDisjointSetUnion(n)

	cost := 0
	edgesUsed := 0

	// Step 4: Kruskal
	for _, edge := range edges {
		if dsu.Union(edge.u, edge.v) {
			cost += edge.d
			edgesUsed++

			// MST is complete
			if edgesUsed == n-1 {
				break
			}
		}
	}

	return cost
}