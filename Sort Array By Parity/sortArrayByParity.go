func sortArrayByParity(A []int) []int {
    for left, right := 0, 0; right < len(A); right++ {
        if A[right] % 2 == 0 {
            A[left], A[right] = A[right], A[left]
            left++
        }
    }
    return A
}
