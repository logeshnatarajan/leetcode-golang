func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    i,j,count:=0,0,0
    for (i<len(g)&&j<len(s)){
        if g[i]<=s[j]{
            i+=1
            j+=1
            count+=1
            continue
        }else{
            j+=1
        }
    }
    return count
}
