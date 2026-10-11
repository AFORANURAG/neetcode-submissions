func getConcatenation(nums []int) []int {
  n:=len(nums)
  arr:= make([]int,2*n)

  for i:=0;i<2*n;i++{
    idx:=(i)%n
    arr[i]=nums[idx]
  }
  return arr
}
