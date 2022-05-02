func isSubsequence(s string, t string) bool {
    i,j:=0,0
    for ;j<len(t) && i<len(s) ;{
        if s[i]==t[j]{
            i++ 
            j++
            continue
        }
        j++
    }

    if i==len(s) {
        return true
    }
    return false
}
