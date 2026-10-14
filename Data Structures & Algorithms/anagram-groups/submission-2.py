class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # I am solve it using sorting
        # brute force will take upto n^3 time complexity and n2 space complexity

        m = {}
        # key-->[]
        for s in strs:
            # s here is a string
            sorted_s =  "".join(sorted(s))
            if sorted_s not in m:
                m[sorted_s] = [s]
            else:
                m[sorted_s] = m[sorted_s]+[s]

        ans = [v for k,v in m.items()]
        return ans