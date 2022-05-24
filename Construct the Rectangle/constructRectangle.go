func constructRectangle(area int) []int {
    sq:=int(math.Sqrt(float64(area)))
    fmt.Println(math.Sqrt(99))
    fmt.Println(int(math.Sqrt(99)))
    for {
        temp:=area/sq
        if (temp*sq)==area{
            if temp>sq{
                return []int{temp,sq}
            }else{
                return []int{sq,temp}
            }
        }
        sq--
    }
}
