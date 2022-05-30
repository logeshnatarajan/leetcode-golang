func fib(n int) int {
    a,b,res:=0,1,0
    if n==1{
        return b
    }
    for i:=1;i<n;i++{
        res=a+b
        a=b
        b=res
    }
    return res
}
