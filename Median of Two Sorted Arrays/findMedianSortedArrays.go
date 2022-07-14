func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    compared := 0
    total := len(nums1) + len(nums2)
    middle := float64(total) / 2 + 0.5
    nums1Cur := 0
    nums2Cur := 0
    prevVal := 0
    currentVal := 0
    
    for float64(compared) < middle {
        prevVal = currentVal
        
        if nums1Cur > len(nums1) - 1 {
            currentVal = nums2[nums2Cur]
            nums2Cur++
        } else if nums2Cur > len(nums2) - 1 {
            currentVal = nums1[nums1Cur]
            nums1Cur++
        } else if nums1[nums1Cur] < nums2[nums2Cur] {
            currentVal = nums1[nums1Cur]
            nums1Cur++
        } else {
            currentVal = nums2[nums2Cur]
            nums2Cur++
        }
        
        compared++
    }
    
    if float64(compared) == middle {
        return float64(currentVal)
    } else {
        return float64(currentVal + prevVal) / 2 
    }
    
    return 0
}
