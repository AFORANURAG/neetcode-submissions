func ladderLength(beginWord string, endWord string, wordList []string) int {
    set:=make(map[string]bool)
	for _,word:=range wordList{
		set[word]=true
	}
	if !set[endWord]{
		return 0
	}
	// we have the set ready
	// lets write that algo to generate combinations
	var generate_neighbours func(string)[]string
	generate_neighbours = func(word string)[]string{
		neighbours:=make([]string,0)
		// for each character, there are 26 possibilities
		for i:=0;i<len(word);i++ {
			// now we have the character
			// we can certainly use utf values so 
			for c:='a';c<='z';c++{
				if c==rune(word[i]){
					continue
				}
				// now lets generate combination
				comb:=word[:i] + string(c) + word[i+1:]
				if set[comb]{
neighbours = append(neighbours,comb)
				}

			}
		}
		return neighbours
	}
	q:=make([]string,0)
	visited:=make(map[string]bool)
	levels:=1
	q = append(q,beginWord)
	visited[beginWord]= true
	for len(q)>0{
		size:=len(q)
        for i:=0;i<size;i++{
curr:=q[0]
		q = q[1:]
		if curr==endWord{
			return levels
		}
		neighs:=generate_neighbours(curr)
		if len(neighs)==0{return 0}
		for _,nei:=range neighs {
			fmt.Println("nei is",nei)
			if !visited[nei]{
				visited[nei]=true
				q = append(q,nei)
			}
		}
		}
		
		levels++
	}
return 0
}
