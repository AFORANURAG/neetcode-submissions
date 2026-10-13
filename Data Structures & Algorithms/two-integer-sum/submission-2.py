class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        # brute forcing is one way,
        # another is using two pointers

        # with hashmap

        m = {}

        for idx,n in enumerate(nums):
           s = target - n
           if s in m and idx!=m[s]:
            return [m[s],idx]
           m[n] = idx








        