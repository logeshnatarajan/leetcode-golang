func islandPerimeter(grid [][]int) int {
    var res int
    dirs := []int{1, 0, -1, 0, 1}
    rln:=len(grid)
    cln:=len(grid[0])
    var find func(i,j int)
    find = func(i,j int){
        if i>=rln || j>=cln ||i<0||j<0 || grid[i][j]==0{
            res++
            return
        }
		if grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		for x := 0; x < 4; x++ {
			xi := i + dirs[x]
			yi := j + dirs[x+1]
			find(xi, yi)
		}
    }

    for i,v:=range grid{
        for j,val:=range v{
            if val==1{
                find(i,j)
                return res
            }
        }
    }
   return res
}
