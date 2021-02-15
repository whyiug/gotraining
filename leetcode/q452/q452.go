package main


import (
	"fmt"
	"sort"
)


func main()  {
	fmt.Println( findMinArrowShots([][]int {{1,2},{2,3},{3,4},{1,3}}))
	fmt.Println( findMinArrowShots([][]int {{1,2},{1,2},{1,2}}))
	fmt.Println( findMinArrowShots([][]int {{1,2},{2,3}}))
}


func findMinArrowShots(points [][]int) int {
	var r int
	r = 0
	len := len(points)
	if len <= 1 {
		return len - r
	}
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	b := 0
	for i:=1; i<len ; i ++ {
		//下一个的开始区间 小于 上一个的结束区间
		if points[i][0] <= points[b][1] {
			r ++
		} else {
			b = i
		}
	}
	return len - r
}