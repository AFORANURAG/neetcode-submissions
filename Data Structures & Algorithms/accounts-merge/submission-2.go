type UnionFind struct {
    parent []int
    rank   []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        rank[i] = 1
    }
    return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
    if x != uf.parent[x] {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
    px, py := uf.Find(x), uf.Find(y)
    if px == py {
        return false
    }
    if uf.rank[px] > uf.rank[py] {
        uf.parent[py] = px
        uf.rank[px] += uf.rank[py]
    } else {
        uf.parent[px] = py
        uf.rank[py] += uf.rank[px]
    }
    return true
}


func accountsMerge(accounts [][]string) [][]string {
n:=len(accounts)
uf:=NewUnionFind(n)
emailToAcc := make(map[string]int)
for i,account:=range accounts{
	// iterate through all the emails
	for j:=1;j<len(account);j++{
		email := account[j]
		if idx,exists:=emailToAcc[email];exists{
			// union them
			uf.Union(idx,i)
		}else{
			emailToAcc[email] = i
		}
	}
}

// We would have done unions so now accounts having common emails will be together
// for all emails, form groups
emailGroup := make(map[int][]string)

for email,accId:=range emailToAcc{
	// find parent
 leader := uf.Find(accId)
 emailGroup[leader]=append(emailGroup[leader],email)
}
 res := [][]string{}
    for accountId, emails := range emailGroup {
		name:=accounts[accountId][0]
        sort.Strings(emails)
        merged := append([]string{name}, emails...)
        res = append(res, merged)
    }
return res
}

