func getRow(n int) []int {
    var res [][]int
    res=append(res,[]int{1})
    res=append(res,[]int{1,1})
    if n<=1{
        return res[n]
    }
    for i:=2;i<=n;i++{
        temp:=make([]int,i+1)
        temp[0],temp[i]=1,1
        for j:=1;j<i;j++{
            temp[j]=res[i-1][j]+res[i-1][j-1]            
        }
        res=append(res,temp)
    }
    return res[n]
}
