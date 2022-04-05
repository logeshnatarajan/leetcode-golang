func arrangeCoins(n int) int {

    temp,count:=0,0
    for i:=1;i<=n;i++{
        if temp>n{
            return count-1
        }
        temp+=i
        count++
    }
    fmt.Println(temp,n)
    if temp==n{
        return count
    }else {
        return count-1
    }
}
