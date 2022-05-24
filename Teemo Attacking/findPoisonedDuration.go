func findPoisonedDuration(timeSeries []int, duration int) int {
    result,ln:=duration,len(timeSeries)
    for i:=1;i<ln;i++{
        temp:=timeSeries[i]-timeSeries[i-1]
        if temp>=duration{
            result+=duration
        } else{
            result+=temp
        }
    }
    return result
}
