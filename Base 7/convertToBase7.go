func convertToBase7(num int) string {
    if num==0{
        return "0"
    }
    minus:=""
    if num<0{
        num*=-1
        minus="-"
    }
    res:=""
    for num!=0{
        res=strconv.Itoa(num%7)+res
        num/=7
    }
    return minus+res
}
