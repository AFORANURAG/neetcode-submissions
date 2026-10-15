class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        # I mean, its does not have to be longest but anything that is same across all the strings
        # Common prefix, that is it, we need to find common prefix
        # pick any string, then pick its character, and then check if its common across all the other strings

        # find the shortest string in here

        n = len(strs)
        
        if n == 1:
            return strs[0] 
        
    
        val = strs[0]
        shortest_length = len(val)
        print("val is",val)
        while shortest_length > 0:
            all_common = True
            print("val is",val[0:shortest_length])
            print("shortest length is",shortest_length)
            for idx,s in enumerate(strs):
                if idx==0:
                    continue
                if not strs[idx].startswith(val[0:shortest_length]):
                    shortest_length = shortest_length - 1
                    all_common = False
            if all_common:
                print("hello")
                return val[0:shortest_length]

        
        return ""
        







        