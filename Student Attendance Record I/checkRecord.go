func checkRecord(s string) bool {
    a,l:=0,0
    for _,v:=range s {
        if v==65{
            a++
        }else if v==76 {
            l++
        }
        if a>=2 || l==3{
            return false
        }
        if v!=76 && l!=0{
            l=0
        }
    }
    return true
}
