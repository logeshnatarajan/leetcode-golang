func peakIndexInMountainArray(A []int) int {
    i:=2
    for ; i < len(A); i++ {
        if A[i] < A[i-1] {
            break
        }
    }
    return i-1
}
