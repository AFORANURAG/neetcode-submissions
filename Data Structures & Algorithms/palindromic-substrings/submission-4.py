class Solution:
    def countSubstrings(self, s: str) -> int:
        # find brute force solution first
        # generate all substring first, then check how many of them are palindromes.
        # a,b,c
        # ab,ac
        # bc
        # memo
        n = len(s)
        max_length = 0
        k = 0
        count = 0
        for i in range(n):
            l,r = i,i

            while l>=0 and r<n and s[l]==s[r]:
                if (r-l+1) > max_length:
                    max_length = r- l + 1
                    k = l
                l = l - 1
                r = r + 1
                count = count + 1

                
            l,r = i,i+1
            while l>=0 and r<n and s[l]==s[r]:
                if (r-l+1) > max_length:
                    max_length = r- l + 1
                    k = l
                l = l - 1
                r = r + 1
                count  = count + 1
        print("count is",count)
            
        return count
      

    def isPalindrome(self,s:str) -> bool:
        n = len(s)
        i,j = 0,n-1
        while i<=j:
            if s[i]!=s[j]:
                return False
            i = i + 1
            j = j - 1
        return True




        