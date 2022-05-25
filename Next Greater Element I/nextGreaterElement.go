// simple solution 


func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res:=make([]int,len(nums1))
    var m2 = map[int]int{}
    for i,v:=range nums1{
        m2[v]=i
        res[i]=-1
    }
    temp:=[]int{}
    for _,v:=range nums2{
        if len(temp)!=0 && temp[len(temp)-1]<v{
            for len(temp)!=0 && v>temp[len(temp)-1]{
                res[m2[temp[len(temp)-1]]]=v
                temp=temp[:len(temp)-1]
            }
            
        }
        if _, ok:=m2[v];ok{
            if len(temp)==0{
                temp=append(temp,v)
            }else if temp[len(temp)-1]>v{
                    temp=append(temp,v)
            }

        }

    }
    return res
}
