type DSU struct {
	parent []int
	rank   []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU{
		parent: parent,
		rank:   rank,
	}
}

func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x])
	}

	return dsu.parent[x]
}

func (dsu *DSU) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)

	// Already connected -> adding this edge creates a cycle
	if rootX == rootY {
		return false
	}

	// Attach smaller-rank tree under larger-rank tree
	if dsu.rank[rootX] < dsu.rank[rootY] {
		dsu.parent[rootX] = rootY
	} else if dsu.rank[rootX] > dsu.rank[rootY] {
		dsu.parent[rootY] = rootX
	} else {
		// Same rank
		dsu.parent[rootY] = rootX
		dsu.rank[rootX]++
	}

	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	// Add original index:
	// [u, v, weight, originalIndex]
	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	// Sort by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// --------------------------------------------------
	// Step 1: Construct the normal MST
	// --------------------------------------------------

	dsu := newDSU(n)

	mstWeight := 0
	edgesUsed := 0

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]

		if dsu.Union(u, v) {
			mstWeight += w
			edgesUsed++
		}
	}

	criticalEdges := make([]int, 0)
	pseudoCriticalEdges := make([]int, 0)

	// --------------------------------------------------
	// Step 2: Test every edge
	// --------------------------------------------------

	for i := 0; i < len(edges); i++ {

		// ==================================================
		// Test whether edge i is CRITICAL
		// ==================================================

		dsu = newDSU(n)

		mstw := 0
		edgesUsed := 0

		for j := 0; j < len(edges); j++ {

			// Exclude edge i
			if i == j {
				continue
			}

			e := edges[j]

			u, v, w := e[0], e[1], e[2]

			if dsu.Union(u, v) {
				mstw += w
				edgesUsed++
			}
		}

		// If we cannot construct a spanning tree,
		// OR the resulting MST is more expensive,
		// edge i is critical.
		if edgesUsed != n-1 || mstw > mstWeight {
			criticalEdges = append(
				criticalEdges,
				edges[i][3], // original index
			)

			// No need to test pseudo-critical
			// because a critical edge cannot be pseudo-critical.
			continue
		}

		// ==================================================
		// Test whether edge i is PSEUDO-CRITICAL
		// ==================================================

		dsu = newDSU(n)

		e := edges[i]

		u, v, w := e[0], e[1], e[2]

		// FORCE edge i into the MST
		dsu.Union(u, v)

		pmstw := w
		edgesUsed = 1

		// Run Kruskal with the forced edge already included
		for j := 0; j < len(edges); j++ {

			if i == j {
				continue
			}

			e := edges[j]

			u, v, w := e[0], e[1], e[2]

			if dsu.Union(u, v) {
				pmstw += w
				edgesUsed++
			}
		}

		// If forcing this edge still produces
		// an MST with the same minimum weight,
		// it is pseudo-critical.
		if edgesUsed == n-1 && pmstw == mstWeight {
			pseudoCriticalEdges = append(
				pseudoCriticalEdges,
				edges[i][3], // original index
			)
		}
	}

	return [][]int{
		criticalEdges,
		pseudoCriticalEdges,
	}
}