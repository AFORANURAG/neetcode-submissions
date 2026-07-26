import (
    "slices"
)
func findItinerary(tickets [][]string) []string {
 // We will start from JFK and then start that Hierholzer algo
 // Also its mentioned that the person departs from JFK
 // and we need to keep the below in mind
 // If there are multiple valid flight paths, return the lexicographically smallest one.
 // lets build adjacency list
 adj:=make(map[string][]string)
 for _,ticket:=range tickets{
    // We have a ticket here where ticket[0] is from_i and ticket[1] is to_i
    src,dest:=ticket[0],ticket[1]
    adj[src] = append(adj[src],dest)
  }
  // Now we should also sort such that lexi wise small words appear at last of the array for
  // each node
  // Now we have an adjacency list
  for k,_:=range adj{
  
  sort.Sort(sort.Reverse(sort.StringSlice(adj[k])))
  fmt.Println("after sorting k is",k)
  fmt.Println("after sorting adj[k] is",adj[k])
  }

  itenarary:=make([]string,0)
  var dfs func(node string)
  dfs = func(node string){
    for len(adj[node])>0{
        // So adj[node] is essentially an array
        l:=len(adj[node])
        lastNode:=adj[node][l-1]
        adj[node]=adj[node][:l-1]
        dfs(lastNode)
    }
    itenarary = append(itenarary,node)

  }
  dfs("JFK")
  fmt.Println("itenarary is",itenarary)
  slices.Reverse(itenarary)

  return itenarary
}