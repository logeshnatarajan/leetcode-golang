func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1)>len(nums2){
        nums1,nums2=nums2,nums1
    }
    m:=make(map[int]bool)
    dm:=make(map[int]bool)
    for _,v:=range nums1{
        m[v]=true
    }
    res:=[]int{}
    for _,v:=range nums2{
        if m[v] && !dm[v]{
            res=append(res,v)
        }
        dm[v]=true
    }
    return res
}
