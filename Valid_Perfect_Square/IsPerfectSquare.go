func isPerfectSquare(num int) bool {
    if num<=1000000{
        i:=1
        for {
            temp:=i*i
            if temp==num{
                return true
            }else if temp>num{
                return false
            }
            i++
        }
    }else{
        i:=1000
        for {
            temp:=i*i
            if temp==num{
                return true
            }else if temp>num{
                return false
            }
            i++
        }
    }
}
