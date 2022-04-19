type NumArray struct {
    v []int
    sums []int
}


func Constructor(nums []int) NumArray {
    temp:=make([]int,len(nums))
    temp[0]=nums[0]
    for i:=1;i<len(nums);i++{
        temp[i]=temp[i-1]+nums[i]
    }
    return NumArray{nums,temp}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left==0{
        return this.sums[right]
    }
    return this.sums[right]-this.sums[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
