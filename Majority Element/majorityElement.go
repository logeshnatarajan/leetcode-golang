func majorityElement(nums []int) int {
    num,max:=0,0
    for _,v:=range nums{
        if num==0{
            max=v
        }
        if max==v{
            num++
        }else{
            num--
        }        
    }
    return max
}
