func hammingDistance(x int, y int) int {
    xBin := fmt.Sprintf("%031b", x)
    yBin := fmt.Sprintf("%031b", y)
    ret := 0
    for i:=0;i<31;i++ {
        if xBin[i] != yBin[i] {
            ret++
        }
    }
    return ret

}
