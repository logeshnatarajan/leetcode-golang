func countNegatives(grid [][]int) int {
    var res int
    for _,val:=range grid{
        left,right:=0,len(val)-1
        for left<=right{
            mid := left + (right - left) / 2
            if val[mid]<0{
                right=mid-1
            }else{
                left=mid+1
            }
        }
        res+=(len(val)-left)
    }
    
    return res
}
