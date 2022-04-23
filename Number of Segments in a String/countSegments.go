func countSegments(s string) int {
    segs := strings.Split(s, " ")
    count := 0
    for _, i := range segs {
        if len(i) != 0 {
            count ++
        }
    }
    return count
}
