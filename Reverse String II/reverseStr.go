func reverseStr(s string, k int) string { 
    tmp := []byte(s)
    end := len(tmp)-1 
    for i := 0; i <= end; i += 2*k {
        x := i
        y := i + k - 1
        if y > end {
            y = end
        }
        for x < y {
            tmp[x], tmp[y] = tmp[y], tmp[x]
            x++
            y--
        }
    } 
    return string(tmp)    
}
