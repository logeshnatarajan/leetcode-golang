func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := float64(sum)
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		max = math.Max(max, float64(sum))
	}
	return max / float64(k)
}

// another solution 

func max(a,b float64)float64{
    if a > b{
        return a
    }
    return b
}
func findMaxAverage(nums []int, k int) float64 {
    maxAverage := float64(math.MinInt32)
    sum := 0
    for i:=0; i<k;i++{
        sum += nums[i]
    }
    maxAverage = max(maxAverage,float64(sum)/float64(k))
    for i:=k;i<len(nums);i++{
        sum += nums[i]
        sum -= nums[i-k]
        maxAverage = max(maxAverage,float64(sum)/float64(k))
    }
    
    return maxAverage
}
