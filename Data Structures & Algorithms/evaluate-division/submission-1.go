func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    // adj_list first
    adj := make(map[string][][2]interface{})
    for i,eq:=range equations{
        a,b:=eq[0],eq[1]
         adj[a] = append(adj[a], [2]interface{}{b, values[i]})
        adj[b] = append(adj[b], [2]interface{}{a, 1.0 / values[i]})
    }
    // lets move on to next steps
    var bfs func(string,string)float64

    bfs = func(src,target string)float64{
        // if src or target not present
        if _,exists:=adj[src];!exists{
            return -1.0
        }
        if _,exists:=adj[target];!exists{
            return -1.0
        }
        queue := [][2]interface{}{{src, 1.0}}
        visited := make(map[string]bool)
        visited[src]=true

        for len(queue)>0{
            // pop
            curr:=queue[0][0].(string)
            w:=queue[0][1].(float64)
            queue = queue[1:]
            if curr==target{
                return w
            }
            // visit all the neibhours
            for _,nei:=range adj[curr]{
                neiNode := nei[0].(string)
                neiWeight := nei[1].(float64)
                if !visited[neiNode]{
                    visited[neiNode]=true
                    queue = append(queue,[2]interface{}{neiNode,neiWeight*w})
                }
            }
        }    
        
 return -1.0
    }

      res := make([]float64, len(queries))
    for i, q := range queries {
        res[i] = bfs(q[0], q[1])
    }

    return res
}