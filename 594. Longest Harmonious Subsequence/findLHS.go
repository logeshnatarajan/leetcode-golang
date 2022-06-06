func findLHS(nums []int) int {
  sort.Ints(nums)
  res := 0
  i, j := 0, 0
  for j < len(nums) {
    if nums[j]-nums[i] > 1 {
      for nums[j] - nums[i] > 1 {
        i++
      }
    } else if nums[i] != nums[j] {
      res = max(res, j-i+1)
    }
    j++
  }
  return res
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
