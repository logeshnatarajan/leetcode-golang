func longestPalindrome(s string) string {
	result := ""
	if len(s)%2 == 1 {
		for i := 0; i < len(s); i++ {
			// odd lenght
			l, r := i, i
			call(l, r, &result, s)

		}
	} else {
		for i := 0; i < len(s); i++ {

			// even lenght
            l, r := i, i+1
			call(l, r, &result, s)

		}
	}

	return result
}
func call(l, r int, res *string, s string) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if len(*res) < (r - l + 1) {
			*res = s[l : r+1]
		}
		l--
		r++
	}
}
