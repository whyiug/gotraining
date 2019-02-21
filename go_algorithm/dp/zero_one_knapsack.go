package main

import "fmt"

/**
有 N 种物品，第 i 种物品的重量为 wi，价值为 vi，每种物品只有一个。背包能承受的重量为 W 。将哪些物品装入背包可使这些物品的总重量不超过背包容量，且价值总和最大?
*/
func main() {
	// N 种物品
	n := 3
	// W 总容量
	w := 10
	// 所有物品重量
	ww := []int{0, 3, 4, 5}
	// 所有物品价值
	vv := []int{0, 4, 5, 6}
	fmt.Println(dp(n, w, ww, vv))
}

func dp(n int, w int, ww []int, vv []int) int {
	var dp [4][11]int
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			// dp[i][j] = max(dp[i-1][j], dp[i-1][j-ww[i]] + vv[i])
			if j >= ww[i] {
				if dp[i-1][j] > dp[i-1][j-ww[i]]+vv[i] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j-ww[i]] + vv[i]
				}
			}
		}
	}
	// fmt.Println(dp)
	vmax := dp[n][w]
	return vmax
}
