func findDisappearedNumbers(nums []int) []int {
    m:=make(map[int]bool)
    for _,v:=range nums {
        m[v]=true
    }
    res:=[]int{}    
    for i:=1;i<=len(nums);i++{
        if !m[i]{
            res=append(res,i)
        }
    }
    return res
}
