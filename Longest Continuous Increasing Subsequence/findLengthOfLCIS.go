func findLengthOfLCIS(nums []int) int {
    ln,max,temp:=len(nums),1,1
    for l,r:=0,1;r<ln;r,l=r+1,l+1 {
        if nums[r]>nums[l]{
            temp++
        }else{
            temp=1
        }
        if max<temp{
            max=temp
        }
    }
    return max
}
