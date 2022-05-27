func findWords(words []string) []string {
    top:="qwertyuiop"
    top+=strings.ToUpper(top)
    middle:="asdfghjkl"
    middle+=strings.ToUpper(middle)
    last:="zxcvbnm"
    last+=strings.ToUpper(last)
    var m1,m2,m3=map[uint8]int{},map[uint8]int{},map[uint8]int{}
    for i:=0;i<20;i++{
        m1[top[i]]=1
        if i<18{
            m2[middle[i]]=2
        }
        if i<14{
            m3[last[i]]=3
        }
    }
    var res []string
    for _,str:=range words {
        found:=true
        if _,ok:=m1[str[0]];ok{
            for _,vv:=range str {
                if _, k :=m1[uint8(vv)];!k{
                    found=false
                    break
                }
            }
        }else if _,ok:=m2[str[0]];ok{
            for _,vv:=range str {
                if _, k :=m2[uint8(vv)];!k{
                    found=false
                    break
                }
            }
        }else if _,ok:=m3[str[0]];ok{
            for _,vv:=range str {
                if _, k :=m3[uint8(vv)];!k{
                    found=false
                    break
                }
            }
        }
        if found {
            res=append(res,str)
        }
    }
    return res
}
