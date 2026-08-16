func foreignDictionary(words []string) string {
    adj := make(map[rune]map[rune]struct{})
    for _, w := range words {
        for _, c := range w {
            if _, exists := adj[c]; !exists {
                adj[c] = make(map[rune]struct{})
            }
        }
    }
	

	// lets iterate over the words

	for i:=0;i<len(words)-1;i++{
		w1,w2:=words[i],words[i+1]
		minLen := len(w1)
        if len(w2) < minLen {
            minLen = len(w2)
        }
		 if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
            return ""
        }
		for j:=0;j<minLen;j++{
		if w1[j]!=w2[j]{
		adj[rune(w1[j])][rune(w2[j])]=struct{}{}
		break
		}
		}
	}

	// we require a visit
	visited := make(map[rune]int)
	res:=make([]rune,0)
	var dfs func(rune)bool

	dfs = func(char rune)bool{
		if val,ok:=visited[char];ok{
			return val==1
		}
		// this adj[char] returns a map {} and the iterator takes it and gives us values in nei
		// 1:{
//	}

visited[char] = 1
		for nei:=range adj[char]{
			// adj[char] returns you a map
			if dfs(nei){
				return true
			}
		}
		visited[char]=-1
		res = append(res,char)
		return false
	}

	for char:=range adj{
		if dfs(char){return ""}
	}
fmt.Println("res is",res)
	ans:=make([]byte,0)
	for i:=len(res)-1;i>=0;i--{
		ans = append(ans,byte(res[i]))
	}

return string(ans)
}

// It takes two words and returns the length of shorter word

func shorterWord(w1,w2 string)int{
	if len(w1)<len(w2){
		return len(w1)
	}
	return len(w2)
}