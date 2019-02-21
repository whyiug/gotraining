package main
/**
golang参数传递，一定为值传递
map, slice, chan被当做指针处理
https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
 */
import "fmt"

func main33() {
	// 切片作为函数参数
	a := []struct{
		itemId int
		name string
	}{
		{1, "hello"},
		{2, "world"},
	}
	fmt.Printf("%+v\n", a)
	modify22(a)
	fmt.Printf("%+v\n", a)
	return

	// 切片作为函数参数
	i := []int{1, 2, 3}
	fmt.Printf("%v\n", i)
	modify2(i)
	fmt.Printf("%v\n", i)

	// 数组作为函数参数
	m := [3]int{1, 2, 3}
	fmt.Printf("%v\n", m)
	modify3(&m)
	fmt.Printf("%v\n", m)

	// 指针作为函数参数
	n := 1
	np := &n
	fmt.Printf("np的值为 %v\n", np)
	fmt.Printf("np的地址（指针）为 %v\n", &np)
	modify4(np)
}

func modify22(s []struct{
	itemId int
	name string
})  {
	s[0].itemId = 3
	return
}

func modify2(ia []int)  {
	ia[0] = 11
	return
}

func modify3(ia *[3]int)  {
	(*ia)[0] = 11
	return
}

func modify4(ia *int)  {
	fmt.Printf("ia的值为 %v\n", ia)
	fmt.Printf("ia的地址（指针）为 %v\n", &ia)
	*ia = 11
	return
}
