func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var result int
	charIndexMap := make(map[uint8]int)
	var start int
	for end := 0; end < n; end++ {
        if duplicateIndex, isDuplicate := charIndexMap[s[end]]; isDuplicate{
			result = max(result, end-start)
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}
			start = duplicateIndex + 1
        }
		charIndexMap[s[end]] = end
	}
	result = max(result, n-start)
	return result
}
func max(n,m int)int{
    if n>m{
        return n
    }else{
        return m
    }
}

