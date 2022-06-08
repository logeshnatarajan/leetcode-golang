func findRestaurant(list1 []string, list2 []string) []string {
    var result []string
    var m1 = map[string]int{}
    for i,v:=range list1{
        m1[v]=i
    }
    min:=2000
    for i,v:=range list2{
        if num,ok:=m1[v];ok{
            if (num+i)==min{
                result=append(result,v)
            }else if (num+i)<min{
                min=num+i
                result=[]string{}
                result=append(result,v)
            }
        }
    }
    return result
}
