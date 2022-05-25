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
 // efficient solution 

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nums2Map := map[int]int{}
    
    for i, num := range nums2 {
        nums2Map[num] = i
    }
    
    result := make([]int, len(nums1))
    for i, num := range nums1 {
        found := false
        for nextIndex := nums2Map[num] + 1; nextIndex < len(nums2); nextIndex++ {
            if nums2[nextIndex] > num {
                result[i] = nums2[nextIndex]
                found = true
                break
            }
        }
        
        if !found {
            result[i] = -1
        }
    }
    
    return result
}
