
// -------------------- Graph --------------------

type Flight struct {
	Flight int
	Price  int
}

// -------------------- Priority Queue --------------------

type FlightPriorityQueueItem struct {
	Vertex   int
	Distance int
	Stops    int
}

type FlightPriorityQueue []*FlightPriorityQueueItem

func (pq FlightPriorityQueue) Len() int {
	return len(pq)
}

func (pq FlightPriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance // Min Heap
}

func (pq FlightPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *FlightPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*FlightPriorityQueueItem))
}

func (pq *FlightPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil // avoid memory leak

	*pq = old[:n-1]

	return item
}

// implemented priority queue
// now we will implement a flavored version of the priority queue
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]Flight, n)

	for _, flight := range flights {
		from_i, to_i, price := flight[0], flight[1], flight[2]
		graph[from_i] = append(graph[from_i], Flight{
			Flight: to_i,
			Price:  price,
		})
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, k+2)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	pq := &FlightPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &FlightPriorityQueueItem{
		Vertex:   src,
		Distance: 0,
		Stops:    0,
	})
	dist[src][0] = 0
	// now its gonna be very simple and easy
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*FlightPriorityQueueItem)
		flight, d, stops := item.Vertex, item.Distance, item.Stops

		if flight == dst {
			return d
		}
		if stops>=k+1 || dist[flight][stops] < d {
			continue
		}
		for _, neiFlight := range graph[flight] {
			neiFlight, neiFlightPrice := neiFlight.Flight, neiFlight.Price
			newDistance := d + neiFlightPrice
			newStops := stops + 1
			if newDistance < dist[neiFlight][newStops] {
				dist[neiFlight][newStops] = newDistance
				heap.Push(pq, &FlightPriorityQueueItem{
					Vertex:   neiFlight,
					Distance: newDistance,
					Stops:    newStops,
				})
			}
		}

	}
   
    return -1
}
