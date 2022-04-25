func thirdMax(nums []int) int {
    sort.Ints(nums)
    num,temp,count:=0,0,2147483649
    for i:=len(nums)-1;i>=0;i--{
        temp=nums[i]
        if count>temp{
            count=temp
            num++
        }
        if num==3{
            return temp
        }
    }
    return nums[len(nums)-1]
}
