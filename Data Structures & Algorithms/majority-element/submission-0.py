class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        # sorting is not a sure shot way
        # {5:4,1:3}
        # [5,5,1,1,1,5,5] original array
        # sort in place, [1,1,1,5,5,5,5]
        n = len(nums)
        nums.sort()
        return nums[n//2]
        

        