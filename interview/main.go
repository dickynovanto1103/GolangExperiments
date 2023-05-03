package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func findMiddleNode(head *ListNode) (middle *ListNode, prev *ListNode) {
	fast := head
	slow := head

	prev = head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}

		prev = slow
		slow = slow.Next
	}

	return slow, prev
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	middle, prev := findMiddleNode(head)
	prev.Next = middle.Next

	return head
}

func main() {
	//1,2,5,7,1,2,6
	newList := deleteMiddle(&ListNode{1, &ListNode{2, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}})
	for newList != nil {
		fmt.Printf("%v ", newList.Val)
		newList = newList.Next
	}
	fmt.Println()
	//1,2,5,1,2,6

	fmt.Printf("second case\n")
	newList = deleteMiddle(nil)
	for newList != nil {
		fmt.Printf("%v ", newList.Val)
		newList = newList.Next
	}
	fmt.Println()

	fmt.Printf("third case\n")
	newList = deleteMiddle(&ListNode{1, nil})
	for newList != nil {
		fmt.Printf("%v ", newList.Val)
		newList = newList.Next
	}
	fmt.Println()

}
