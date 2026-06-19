func combinationSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result:=make([][]int,0)
    curr:=make([]int,0)

    var dfs func(int, []int, int)
    dfs = func (i int,curr []int, total int){
        if total==target{
            temp:=make([]int,len(curr))
            copy(temp,curr)
            result=append(result,temp)
            return 
        }
        for idx:=i;idx<len(nums);idx++{
            if total+nums[idx]>target{
                return
            }
            curr=append(curr,nums[idx])
            dfs(idx,curr,total+nums[idx])
            curr=curr[:len(curr)-1]
        }
    }
    dfs(0,curr,0)


    return result 
}
