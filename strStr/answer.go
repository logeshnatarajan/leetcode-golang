func strStr(haystack string, needle string) int {
    if strings.Contains(haystack,needle){
        if len(haystack)==0{
            return 0
        }else if (haystack == needle )|| (needle == ""){
            return 0
        }else {
            return(strings.Index(haystack,needle))
        }
    }else{
        return -1
    }
}
