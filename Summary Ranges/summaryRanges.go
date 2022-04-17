func summaryRanges(nums []int) []string {
    var x []string
    if len(nums)==0{
        return []string{}
    }
    prev,first:=nums[0],nums[0]
    if len(nums)==1{
        return []string{strconv.Itoa(nums[0])}
    }
    for i:=1;i<len(nums);i++{
        if prev+1==nums[i]{
            prev++
        }else{
            if prev==first{
                x=append(x,strconv.Itoa(nums[i-1]))
                prev,first=nums[i],nums[i]
                continue
            }else{
                str:=fmt.Sprintf("%d->%d",first,prev)
                x=append(x,str)
                prev,first=nums[i],nums[i]
                continue
            }
        }
    }
    if prev==first{
        x=append(x,strconv.Itoa(prev))
        return x
    }else{
        str:=fmt.Sprintf("%d->%d",first,prev)
        x=append(x,str)
        return x
    }
    
}
