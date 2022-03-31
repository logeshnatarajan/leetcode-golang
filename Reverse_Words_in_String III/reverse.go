func reverseWords(s string) string {
    res:=strings.Fields(s)
    for i,v:=range res{
        res[i]=rev(v)
    }
    return strings.Join(res," ")
}
func rev(s string)string{
    bs:=[]byte(s)
    ln:=len(bs)
    for i:=0;i<ln/2;i++{
        bs[i],bs[ln-1-i]=bs[ln-1-i],bs[i]
    }
    return string(bs)
}
