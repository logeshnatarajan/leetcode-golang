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

// second soluton which is vere simple but not efficient but use this example to understand

func repeatedSubstringPattern(s string) bool{
    for i:=1;i<=len(s)/2;i++{
        v:=s[:i]
        if s == (strings.Repeat(v,len(s)/(i))){
            return true
        }
    
    }
    return false
}
