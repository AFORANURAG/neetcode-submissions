# class Solution:
#     def coinChange(self, coins: List[int], amount: int) -> int:
#         # say we have [10,5,1]
#         # say
#         # decision tree
#         # 10,10 oops
#         # 10,5, oops
#         # 10,1,1 works
#         # We found a way
#         # now next
#         # 5, no need to check for last or already tried element
#         # 5,5,5-->oops
#         # 5,5,1,1, works--> we found a way
#         # lets move next
#         # 1,1,1,1,1,1,1,1,1,1,1,1, we found a fucking way to arrange.
#         # lets try to first do it using recursion
#         if amount == 0:
#             return 0

#         memo = {}
#         n = len(coins)

#         def dfs(i, a):
#             if a == 0:
#                 return 0

#             if (i, a) in memo:
#                 return memo[(i, a)]

#             min_steps = float('inf')

#             for j in range(i, n):
#                 if a >= coins[j]:
#                     result = dfs(j, a - coins[j])

#                     if result != float('inf'):
#                         min_steps = min(min_steps, 1 + result)

#             memo[(i, a)] = min_steps
#             return min_steps

#         result = dfs(0, amount)

#         return -1 if result == float('inf') else result
        
#         # We have an item

                
class Solution:
    def coinChange(self,coins: [int], amount: int) -> int:
        # say we have [10,5,1]
        # say
        # decision tree
        # 10,10 oops
        # 10,5, oops
        # 10,1,1 works
        # We found a way
        # now next
        # 5, no need to check for last or already tried element
        # 5,5,5-->oops
        # 5,5,1,1, works--> we found a way
        # let's move next
        # 1,1,1,1,1,1,1,1,1,1,1,1, we found a fucking way to arrange.
        # Let's try to first do it using recursion
        if amount == 0:
            return 0
        n = len(coins)
        dp = [amount + 1] * (amount + 1)
        dp[0] = 0

        print("dp array is",dp)

  
        for a in range(1,amount+1):
            # print("i is",i)
            for c in coins:
                if a-c >=0:
                   dp[a] = min(dp[a],1+ dp[a-c])
        print("dp is",dp)


        return dp[amount] if dp[amount]!=amount+1 else -1
        
