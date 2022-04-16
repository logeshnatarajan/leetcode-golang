func generate(num int) [][]int {
    var res [][]int
    res=append(res,[]int{1})
    if num==1{
        return res
    }
    res=append(res,[]int{1,1})
    if num==2{
        return res
    } 
    for i:=2;i<num;i++{
        temp:=make([]int,i+1)
        for j:=0;j<=i;j++{
            if j==0{
                temp[j]=1
                continue
            }
            if j==i{
                temp[j]=1
                continue
            } 
            temp[j]=res[i-1][j]+res[i-1][j-1]
            
        }
        res=append(res,temp)
    }
    return res
}
