func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum, sqrt := 1, int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i != 0 {
			continue
		}
		sum += i
		if i != num/i {
			sum += num / i
		}
	}
	return sum == num
}
