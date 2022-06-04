// simple solution to understand but slightly less efficient 
// speed := 75ms

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    res:=0
    for i:=len(nums)-2;i>=0;i-=2{
        res+=nums[i]
    }
    return res
}


//here the fastest solution 
func arrayPairSum(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    i := 0
    maxSum := 0
    cntSortList, iLb := getCntSortList(&nums)
    
    for true {
        num1 := 0
        num2 := 0
        ok := false

        if ok,i = getNextIndexWithNoneZeroCnt(&cntSortList, i); ok {
            num1 = i + iLb
            cntSortList[i]--
            _,i := getNextIndexWithNoneZeroCnt(&cntSortList, i)
            num2 = i + iLb
            cntSortList[i]--
            maxSum += getMin(num1, num2)
        } else {
            return maxSum
        }
    }    
    return 0
}

func getCntSortList(nums *[]int) ([]int, int) {
    iLb := 10001
    iUb := -10001
    
    for _,num := range(*nums) {
        if num < iLb {
            iLb = num
        }
        
        if num > iUb {
            iUb = num
        }
    }
    
    cntSortList := make([]int, iUb - iLb + 1)

    for _,num := range(*nums) {
        cntSortList[num - iLb]++
    }

    return cntSortList, iLb
}

func getNextIndexWithNoneZeroCnt(cntSortList *[]int, currI int) (bool, int) {
    for i := currI; i < len(*cntSortList); i++ {
        if (*cntSortList)[i] > 0 {
            return true, i
        }
    }

    return false, -1
}

func getMin(num1, num2 int) int {
    if num1 < num2 {
        return num1
    } else {
        return num2
    }
}
