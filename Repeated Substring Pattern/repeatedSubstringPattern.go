func repeatedSubstringPattern(s string) bool {

	for n := 1; n*n <= len(s); n++ {
		if len(s)%n != 0 {
			continue
		}

		if isConstructedSubstr(s, s[0:n]) || isConstructedSubstr(s, s[0:len(s)/n]) {
			return true
		}

	}
	return false
}

func isConstructedSubstr(str string, subStr string) bool {
	if len(str) == len(subStr) {
		return false
	}

	isConstructed := true
	for i := 0; i < len(str); i = i + len(subStr) {
		if str[i:i+len(subStr)] != subStr {
			isConstructed = false
			break
		}
	}
	return isConstructed
}
