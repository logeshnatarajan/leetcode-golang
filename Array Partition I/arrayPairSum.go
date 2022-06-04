// simple solution to understand but slightly less efficient 

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    res:=0
    for i:=len(nums)-2;i>=0;i-=2{
        res+=nums[i]
    }
    return res
}
