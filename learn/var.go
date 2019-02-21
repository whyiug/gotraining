package main

import (
	"fmt"
	"time"
)

func main2() {
	const _x string = "hello world"
	// 声明变量
	var x, y = 3, 4
	var a = 5
	b := 6
	fmt.Println(x, y, a, b, _x)

	var arr [5]int
	fmt.Println(arr)
	arr[1] = 1
	fmt.Println(arr)
	var arr2 = []int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	var arr3 [2][3]int
	for m := 0; m < 2; m++ {
		for n := 0; n < 3; n++ {
			arr3[m][n] = m + n
		}
	}
	fmt.Println("2d: ", arr3)
	j := 1
	for {
		if j > 10 {
			break
		}
		fmt.Println(j)
		j++
	}
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", arr4[2:])
	i := 11
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("invalid value!")
	}
	mapFoo := make(map[string]int)
	mapFoo["one"] = 1
	mapFoo["two"] = 2
	mapFoo["three"] = 3
	fmt.Println(mapFoo)
	delete(mapFoo, "two")
	fmt.Println(mapFoo)
	mapBar := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapBar)
	ii := 1
	pii := &ii
	fmt.Println(ii, pii, *pii)
	ii = 3
	fmt.Println(ii, pii, *pii)
	*pii = 4
	fmt.Println(ii, pii, *pii)
	// ????
	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}
	type Person struct {
		name  string
		age   int
		email string
	}
	person := Person{"tom", 30, "whiyi"}
	fmt.Println(person)
	pPerson := &person
	pPerson.age = 40
	fmt.Println("pPersion", pPerson)
	fmt.Println("persion", person)
	//
	var p = new([]int)
	var v = make([]int, 10)
	fmt.Println(p)
	fmt.Println(v)
	fmt.Println(Max(3, 2))
	mm := map[string]int{"one": 0, "two": 2, "three": 3}
	one, oneOk := mm["one"]
	four, fourOk := mm["four"]
	fmt.Println(one, oneOk)
	fmt.Println(four, fourOk)
	if value, ok := mm["one"]; ok {
		fmt.Println("OK", value)
	} else {
		fmt.Println("false", value)
	}
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
	nextNumFuncs := next()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFuncs())
	}
	fmt.Println(fmt.Sprintf("%d", time.Now().Unix()))
	var feeds = []int64{1, 2, 3, 4}
	iii := 2
	feeds = append(feeds[:iii+1], feeds[iii:]...)
	feeds[iii] = 22
	fmt.Println(feeds)
	fmt.Printf("%T\n", feeds)
	fmt.Printf("%+v\n", feeds)
	fmt.Printf("%#v\n", feeds)
	//
	var ns = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(ns))
}

func maxSubArray(nums []int) int {
	var d [100000]int
	var max int
	for k, v := range nums {
		if k == 0 {
			d[k] = v
			max = v
		} else {
			if d[k-1] > 0 {
				d[k] = d[k-1] + v
			} else {
				d[k] = v
			}
			max = getMax(max, d[k])
		}
	}
	return max
}

func getMax(a int, b int) int {
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func multiRet(key string) (int, bool) {
	mm := map[string]int{"one": 0, "two": 2, "three": 3}
	var val int
	var ok bool
	val, ok = mm[key]
	return val, ok
}

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func next() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}
