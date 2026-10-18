class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:

        # [2,3,4,5,3,2,1,3,7]
        # [2,3,4,5,3,2,1,3,7]
        # [2,4,3,5,3,2,1,3,7]
        # [2,4,5,3,3,2,1,3,7]
        # [2,4,5,3,3,2,1,3,7]
        # [2,4,5,2,3,3,1,3,7]
        # [2,4,5,2,1,3,3,3,7]
        # [2,3,5,2,7,3,3,3,3]

        i,j = 0,1
        l = len(nums)
        occurences = 0

        while i < l and j < l:
            if nums[i]==val and nums[j]!=val:
                nums[i],nums[j] = nums[j],nums[i]
                i = i + 1
                j = j + 1
            elif nums[i]==nums[j]:
                j = j + 1
            else:
                i = i + 1
                j = j + 1
        
        for i in range(l-1,-1,-1):
            if nums[i]==val:
                occurences = occurences + 1

        return l-occurences

