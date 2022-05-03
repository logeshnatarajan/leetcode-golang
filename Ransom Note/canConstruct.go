func canConstruct(r string, m string) bool {
    rl:=len(r)
    ml:=len(m)
    xi:=make([]int,26)
    for i:=0;i<rl || i<ml;i++{
        if  i<rl{
            xi[r[i]-'a']--
        }
        if i<ml{
            xi[m[i]-'a']++
        }
    }
    for _,v:=range xi{
        if v<0{
            return false
        }
    }
    return true
}
