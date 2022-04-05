func findKthPositive(arr []int, k int) int {
    var res []int
    num:=1
    for _,v:=range arr{
        for  ;len(res)<k;num++{
            if num==v{
                num++
                break
            }
            res=append(res,num)
        }
    }
    if len(res)==0 || len(res)<k{
        for i:=arr[len(arr)-1]+1;len(res)<k;i++{
            res=append(res,i)
        }
    }
    return res[k-1]
}
