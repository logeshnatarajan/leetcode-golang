func isValid(s string) bool {
    if len(s)%2==1{
        return false
    }
    var (
        xi []int32
        res bool = true
        m = map[int32]int32{
            41:40,
            93:91,
            125:123,
        }
    )
    
    for _,v:=range s {
        if v==40 || v==91 || v==123 {
            xi=add(xi,v)
            
        }else if v==41 || v==93 || v==125 {
            xi=remove(xi,v,&res,m)
        }
        if !res{
            return false
        }
    }
    if len(xi)==0{
        return true
    }else{
        return false
    }
}
func remove(s []int32,n int32,r *bool,m map[int32]int32)[]int32{
    ln:=len(s)
    if ln==0{
        *r=false
        return s
    }
    if m[n]==s[ln-1]{
        return s[:ln-1]
    }else{
        *r=false
        return s
    }
}
func add(s []int32,n int32)[]int32{
    s=append(s,n)
    return s
}
