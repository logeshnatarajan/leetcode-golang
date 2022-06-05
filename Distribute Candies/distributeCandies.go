func distributeCandies(candyType []int) int {
    cln:=len(candyType)/2
    var m = map[int]bool{}
    for _,v:=range candyType{
        m[v]=true
        if len(m)>=cln{
            return cln
        }
    }
    return len(m)
}
