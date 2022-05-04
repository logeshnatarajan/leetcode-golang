func reverseVowels(ss string) string {
    s:=[]byte(ss)
    var f1,f2 bool
    for i,j:=0,len(s)-1;i<j; {
        if s[i]== 97 || s[i]== 101 || s[i]==105 || s[i]==111 || s[i]==117 || s[i]== 65 || s[i]== 69 || s[i]==73 || s[i]==79 || s[i]==85{
            f1=true
        }else{
            i++
        }
        if s[j]== 97 || s[j]== 101 || s[j]==105 || s[j]==111 || s[j]==117 || s[j]== 65 || s[j]== 69 || s[j]==73 || s[j]==79 || s[j]==85{
            f2=true
        }else{
            j--
        }
        if f1 && f2 {
            f1,f2=false,false
            s[i],s[j]=s[j],s[i]
            i++
            j--
        }
    }
    return string(s)
}
