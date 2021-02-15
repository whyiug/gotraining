package main

import (
	"fmt"
)

func main() {
	head := InitList([]int{3,2,0,-4}, 1)
	fmt.Println(DetectCycle(head))
	fmt.Println(DetectCycle(InitList([]int{1,2}, 0)))
	fmt.Println(DetectCycle(InitList([]int{1}, -1)))

}
//head = [3,2,0,-4], pos = 1
//head = [1,2], pos = 0
//head = [1], pos = -1

 type ListNode struct {
     Val int
     Next *ListNode
 }

func PrintListNodes(head *ListNode) {
	count := 0
	begin := head
	for begin.Next != nil {
		fmt.Println(begin.Val)
		begin = begin.Next
		count ++
		if count > 10 {
			return
		}
	}

}

func InitList(numbers []int, pos int) (head *ListNode) {
	len := len(numbers)
	if len == 0 {
		return nil
	}
	listNodes := make([]ListNode, len)
	for k, number := range numbers {
		listNodes[k] = ListNode{
			Val: number,
			Next: nil,
		}
	}
	for i := 0; i< len-1 ; i ++ {
		listNodes[i].Next = &listNodes[i+1]
	}
	if pos >=0 && pos < len {
		listNodes[len-1].Next = &listNodes[pos]
	}
	head = &listNodes[0]
	return
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result *ListNode
	result = nil
	slow, fast := head, head
	fastStep := 2
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		if fastStep == 2 {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		if slow == fast {
			if slow == head {
				return slow
			} else {
				if fastStep == 2 {
					fast = head
					fastStep = 1
				} else {
					result = fast
					break
				}
			}
		}
	}
	return result
}
