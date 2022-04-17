func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    var xi =[26]int{}
    for i:=0;i<len(s);i++{
        inxi:=s[i]-'a'
        inxt:=t[i]-'a'
        xi[inxi]++
        xi[inxt]--
    }
    for _,v:=range xi {
        if v!=0{
            return false
        }
    }
    return true
}
