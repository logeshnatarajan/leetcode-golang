func longestPalindrome(s string) int {
    m:=make(map[rune]int)
    one,two:=0,0
    for _,v:=range s{
        m[v]+=1
    }
    if len(m)==1{
      return  len(s)
    }
    for _,v:=range m{
        if v%2==0{
            two+=v
            continue
        }else{
            two+=2*(v/2)
            one=1
        }
        
    }
    return two+one
}
