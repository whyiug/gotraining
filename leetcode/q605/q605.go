package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	num := 0
	len := len(flowerbed)
	if len == 1 {
		num = 1 - flowerbed[0]
	} else {
		for i := 0; i < len; i++ {
			if flowerbed[i] == 0 {
				if i == 0 {
					if flowerbed[i+1] == 0 {
						flowerbed[i] = 1
						num++
					}
					continue
				}
				if i == len-1 {
					if flowerbed[i-1] == 0 {
						flowerbed[i] = 1
						num++
					}
					continue
				}
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					num++
				}
			} else {
				continue
			}
		}
	}
	if n <= num {
		return true
	} else {
		return false
	}
}
