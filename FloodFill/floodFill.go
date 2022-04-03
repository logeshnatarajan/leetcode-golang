func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc]==newColor{
        return image
    }
    rln:=len(image)
    cln:=len(image[0])
    source:=image[sr][sc]
    find(rln,cln,source,newColor,sr,sc,image)
    return image
}
func find(r,c,num,newcolor,sr,sc int,image [][]int){
    if (sr<0 || sr>=r || sc<0 || sc>=c){
        return
    }
    if image[sr][sc]!=num{
        return
    }
    image[sr][sc]=newcolor
    find(r,c,num,newcolor,sr+1,sc,image)
    find(r,c,num,newcolor,sr-1,sc,image)
    find(r,c,num,newcolor,sr,sc+1,image)
    find(r,c,num,newcolor,sr,sc-1,image)
    return 
}
