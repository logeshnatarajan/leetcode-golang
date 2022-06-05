func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat)*len(mat[0]) == r*c {
        val:=[]int{}
        for _,v:=range mat{
            for _,vv:=range v{
                val=append(val,vv)
            }
        }
        var res [][]int
        zero:=0
        for i:=0;i<r;i++{
            temp:=[]int{}
            for j:=0;j<c;j++{
                temp=append(temp,val[zero])
                zero++
            }
            res=append(res,temp)
        }
        return res
    }else{
        return mat
    }
}
