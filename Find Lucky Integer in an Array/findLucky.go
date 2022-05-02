func findLucky(arr []int) int {
    m:=make(map[int]int)
    for _,v:=range arr{
        m[v]+=1
    }
    res:=-1
    for k,v:=range m {
        if k==v && v>res{
            res=v
        }
    }
    return res 
}
