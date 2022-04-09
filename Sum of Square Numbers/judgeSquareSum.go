func judgeSquareSum(c int) bool {
    num:=math.Sqrt(float64(c))
    for i,j:=0,int(num);i<=j; {
        if ((i*i)+(j*j))>c{
            j--
        }else if ((i*i)+(j*j))<c{
            i++
        }else{
            return true
        }
    }
    return false
        
}
