package main


import (
	"fmt"
	"sort"
)


func main()  {
	fmt.Println( eraseOverlapIntervals([][]int {{1,2},{2,3},{3,4},{1,3}}))
	fmt.Println( eraseOverlapIntervals([][]int {{1,2},{1,2},{1,2}}))
	fmt.Println( eraseOverlapIntervals([][]int {{1,2},{2,3}}))
}


func eraseOverlapIntervals(intervals [][]int) int {
	var r int
	r = 0
	len := len(intervals)
	if len <= 1 {
		return r
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	b := 0
	for i:=1; i<len ; i ++ {
		//下一个的开始区间 小于 上一个的结束区间
		if intervals[i][0] < intervals[b][1] {
			r ++
		} else {
			b = i
		}
	}
	return r
}