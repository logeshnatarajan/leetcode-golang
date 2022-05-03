func firstUniqChar(s string) int {
    xi:=make([]int,26)
    for i:=0;i<len(s);i++{
        xi[s[i]-'a']+=1
    }
    for i:=0;i<len(s);i++{
        if xi[s[i]-'a']==1{
            return i
        }
    }
    return -1
}
