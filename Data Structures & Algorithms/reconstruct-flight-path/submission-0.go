// func findItinerary(tickets [][]string) []string {
//  // We will start from JFK and then start that Hierholzer algo
//  // Also its mentioned that the person departs from JFK
//  // and we need to keep the below in mind
//  // If there are multiple valid flight paths, return the lexicographically smallest one.

// }


func findItinerary(tickets [][]string) []string {
    adj := make(map[string][]string)

    sort.Slice(tickets, func(i, j int) bool {
        if tickets[i][0] == tickets[j][0] {
            return tickets[i][1] > tickets[j][1]
        }
        return tickets[i][0] > tickets[j][0]
    })

    for _, ticket := range tickets {
        src, dst := ticket[0], ticket[1]
        adj[src] = append(adj[src], dst)
    }

    res := make([]string, 0)

    var dfs func(string)
    dfs = func(src string) {
        for len(adj[src]) > 0 {
            last := len(adj[src]) - 1
            dst := adj[src][last]
            adj[src] = adj[src][:last]
            dfs(dst)
        }
        res = append(res, src)
    }

    dfs("JFK")

    for i := 0; i < len(res)/2; i++ {
        res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
    }

    return res
}