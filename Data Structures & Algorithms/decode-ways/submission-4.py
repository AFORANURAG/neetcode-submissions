class Solution:
    def numDecodings(self,s: str) -> int:
        # What matters is that you put quality reps into it.
        if s=="0":
            return 0
        arr = [i for i in range(26)]
        n = len(s)
        memo = {}
        def dfs(j) -> int:
            count  = 0
            if j==n:
                return 1

            if s[j:n+1] in memo:
                return memo[s[j:n+1]]

            for i in range(j,n):
                matches = re.findall(r"^0\d", s[j:i+1])
                if len(matches) > 0:
                    print("string is",s[j:i+1])
                    print("matches are",matches)
                    return 0
                idx = int(s[j:i+1])    
                if idx>=1  and idx <= len(arr):
                    count = count + dfs(i+1)
            memo[s[j:n+1]]=count
            return count
        return dfs(0)
        