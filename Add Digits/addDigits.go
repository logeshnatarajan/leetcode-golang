func addDigits(num int) int {
    for num>9{
      //  temp:=num
        res:=0
        for num!=0{
            rem:=num%10
            num=num/10
            res+=rem
        }
        num=res
    }
    return num
}
