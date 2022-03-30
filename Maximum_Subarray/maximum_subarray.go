func maxSubArray(nums []int) int {
    res:=nums[0]
    temp:=0
    for _,v:=range nums {
        if temp<0{
            temp=0
        }
        temp+=v
        res=max(temp,res)
    }
    return res
}
func max(a,b int)int{
    if a>b{
        return a 
    }else{
        return b
    }
}
