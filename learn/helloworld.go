package main

import "fmt"

func main()  {
	a := []string{"你好", "2你好", "3你好"}
	fmt.Printf("%v\t%p\n", &a, &a)
	fmt.Printf("%v\t%p\n", &a[0], &a[0])
	fmt.Printf("%v\t%p\n", &a[1], &a[1])
	fmt.Printf("%v\t%p\n", &a[2], &a[2])
}

