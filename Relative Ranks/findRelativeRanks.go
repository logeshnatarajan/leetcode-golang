func findRelativeRanks(score []int) []string {
    temp:= make([]int,len(score))
    m:=copy(temp,score)
    sort.Ints(temp)
    mm:=make(map[int]string)

    for i,sum:=m-1,1;i>=0;i,sum=i-1,sum+1{
        if temp[i]==temp[m-1]{
            mm[temp[i]]="Gold Medal"
            continue
        }else if temp[i]==temp[m-2]{
             mm[temp[i]]="Silver Medal"
            continue
        }else if temp[i]==temp[m-3]{
            mm[temp[i]]="Bronze Medal"
            continue
        }else{
            mm[temp[i]]=strconv.Itoa(sum)
        }

    }
    res:=[]string{}
    for _,v:=range score{
        res=append(res,mm[v])
    }
    return res
}
