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
