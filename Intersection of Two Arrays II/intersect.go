func intersect(nums1 []int, nums2 []int) []int {
    var res []int
    m1:=make(map[int]bool)
    m2:=make(map[int]bool)
        for i,v:=range nums1{
            for ii,vv:=range nums2{
                if v==vv && !m1[i] && !m2[ii]{
                    res=append(res,vv)
                    m1[i]=true
                    m2[ii]=true
                }
            }
        }

    return res
}
