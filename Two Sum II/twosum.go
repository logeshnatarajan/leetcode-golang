func twoSum(numbers []int, target int) []int {
    m:=map[int]int{}
    for inx,v:=range numbers {
        if val, ok:=m[target-v]; ok{
            return []int{val,inx+1}
        }
        m[v]=inx+1
    }
    return []int{}
}

/*  
21 / 21 test cases passed.
Status: Accepted
Runtime: 12 ms
Memory Usage: 5.4 MB
*/
