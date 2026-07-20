// we would require a heap


type Edge struct {
	Vertex int
	Weight int
}
type PriorityQueueItem struct {
	Vertex   int
	Distance int
}
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len()int{
	return len(pq)
}

func (pq PriorityQueue) Less(i,j int)bool{
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i,j int){
	pq[i],pq[j] = pq[j],pq[i]
}

func (pq *PriorityQueue) Push(x interface{}){
	*pq = append(*pq,x.(*PriorityQueueItem))
}

func (pq *PriorityQueue) Pop()interface{} {
	old:=*pq
	n:=len(old)
	item:=old[n-1]
	*pq = old[:n-1]
	old[n-1] = nil
	return item
}

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)

// 	item := old[n-1]
// 	old[n-1] = nil // avoid memory leak

// 	*pq = old[:n-1]

// 	return item
// }

func networkDelayTime(times [][]int, n int, k int) int {
   //standard dijkstra question
   // heap is now implemented
   // lets start implementing
   adj := make([][]Edge, n+1)
   for _,edge:=range times{
	src,dest,w:=edge[0],edge[1],edge[2]
	adj[src]=append(adj[src],Edge{Vertex:dest,Weight:w})
   }
   dist:=make(map[int]int)
   for i:=1;i<=n;i++{
			dist[i] = math.MaxInt

   }
   dist[k] = 0
   pq:=&PriorityQueue{}
   heap.Init(pq)
   heap.Push(pq,&PriorityQueueItem{
	Vertex:k,
	Distance:0,
   })
   lastNode:=n

   for pq.Len() > 0{
	// pop
	curr:=heap.Pop(pq).(*PriorityQueueItem)
	vertex:=curr.Vertex
	d:=curr.Distance
	if dist[vertex]<d{
		continue
	}
	lastNode = vertex
	for _,nei:=range adj[vertex]{
		v:=nei.Vertex 
		w:=nei.Weight
		newDist:= d+w
		if newDist<dist[v]{
			// update the distance
			dist[v]=newDist
			// push to the heap again
			heap.Push(pq,&PriorityQueueItem{Vertex:v,Distance:newDist})
		}
	}
   }
   fmt.Println("arr",dist)

for i:=1;i<=n;i++{
	if dist[i]==math.MaxInt{
		return -1
	}
}
 return dist[lastNode]

}
