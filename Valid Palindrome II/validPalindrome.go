// func validPalindrome(s string) bool {
// 	isPalindrome := func(s string, leftBound, rightBound int) bool {
// 		for left, right := leftBound, rightBound; left < right; left, right = left+1, right-1 {
// 			if s[left] != s[right] {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
// 		if s[left] != s[right] {
// 			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
// 		}
// 	}

// 	return true
// }


func validPalindrome(s string) bool {
    left, right, valid := isPalindrome(s, 0, len(s) - 1)
    if valid {
       return true 
    }
    _, _, leftValid := isPalindrome(s, left, right-1)
    _, _, rightValid := isPalindrome(s, left+1, right)
    return leftValid || rightValid
}

func isPalindrome(s string, left int, right int) (int, int, bool) {
    for left < right {
        if s[left] != s[right] {
            return left, right, false
        }
        left++
        right--
    }
    return left, right, true
}
