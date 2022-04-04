func maxAreaOfIsland(grid [][]int) int {
    r:=len(grid)
    c:=len(grid[0])
    m:=make( map[[2]int]bool)
    var ans func(i,j int)int
    ans=func(i,j int) int{
        if (i<0 || i>=r || j<0 || j>=c || grid[i][j]==0 || m[[2]int{i,j}]){
        return 0
    }
    m[[2]int{i,j}]=true
        return(1+ans(i+1,j)+ans(i-1,j)+ans(i,j+1)+ans(i,j-1))
    }
    var res int
    for i:=0;i<r;i++{
        for j:=0;j<c;j++{
            res=max(res,ans(i,j))
        }
    }
    return res
    
}

func max(a,b int)int{
    if a>b{
        return a
    }else{
        return b
    }
}
