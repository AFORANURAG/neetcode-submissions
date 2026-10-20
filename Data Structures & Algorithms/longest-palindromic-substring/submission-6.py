class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        max_length = 0
        k = 0
        for i in range(n):
            l,r = i,i

            while l>=0 and r<n and s[l]==s[r]:
                if (r-l+1) > max_length:
                    max_length = r- l + 1
                    k = l
                l = l - 1
                r = r + 1

                
            l,r = i,i+1
            while l>=0 and r<n and s[l]==s[r]:
                if (r-l+1) > max_length:
                    max_length = r- l + 1
                    k = l
                l = l - 1
                r = r + 1
        return s[k:k+max_length]
            
