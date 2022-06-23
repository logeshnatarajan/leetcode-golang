func calPoints(ops []string) int {
    stack:=make([]int, 0, 10)
    for i:=0;i<len(ops);i++{
        switch ops[i] {
            case "C":
                stack=stack[:len(stack)-1]
            case "D":
                stack=append(stack, stack[len(stack)-1]*2)
            case "+":
                stack=append(stack, stack[len(stack)-1]+stack[len(stack)-2])
            default:
                val,_:=strconv.Atoi(ops[i])
                stack=append(stack, val)
        }
    }
    ret:=0
    for i:=0;i<len(stack);i++{
        ret+=stack[i]
    }
    return ret
}
