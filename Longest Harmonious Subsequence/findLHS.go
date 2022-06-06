func findLHS(nums []int) int {
    ret := 0
    mp := make(map[int]int, len(nums))   

    for _, n := range nums {
        mp[n]++
    }
    
    for k, v := range mp {
        if mp[k+1] > 0 && v + mp[k+1] > ret {
            ret = v + mp[k+1]
        }
    }
    
    return ret
}
