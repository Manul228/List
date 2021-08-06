package list

import "fmt"

type ListNode struct {
	Val        int
	Next, Prev *ListNode
}

type MyLinkedList struct {
	Head, Tail *ListNode
	Length     int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list.
If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	if index < 0 || index > this.Length-1 {
		return -1
	}

	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

/** Add a node of value val before the first element of the linked list.
After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	item := ListNode{Val: val, Next: nil, Prev: nil}

	if this.Head == nil {
		this.Head = &item
		this.Tail = &item
		this.Length++
	} else {
		item.Next = this.Head
		this.Head.Prev = &item
		this.Head = &item
		this.Length++
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {

	item := ListNode{Val: val, Next: nil, Prev: nil}

	if this.Tail == nil {
		this.Head = &item
		this.Tail = &item
		this.Length++
	} else {
		item.Prev = this.Tail
		this.Tail.Next = &item
		this.Tail = &item
		this.Length++
	}
}

/** Add a node of value val before the index-th node in the linked list.
If index equals to the Length of linked list, the node will be appended to
the end of linked list. If index is greater than the Length, the node will
not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length || index < 0 {
		return
	}

	if index == this.Length {
		this.AddAtTail(val)
		return
	}

	if index == 0 || this.Head == nil {
		this.AddAtHead(val)
		return
	}

	item := ListNode{Val: val, Next: nil, Prev: nil}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	item.Next = current
	item.Prev = current.Prev
	current.Prev.Next = &item
	current.Prev = &item

	this.Length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Length-1 {
		return
	}

	if index == 0 {
		if this.Head.Next != nil {
			this.Head.Next.Prev = nil
		}
		this.Head = this.Head.Next
		this.Length -= 1
		return
	}

	if index == this.Length-1 {
		this.Tail.Prev.Next = nil
		this.Tail = this.Tail.Prev
		this.Length -= 1
		return
	}

	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	current = nil

	this.Length--
}

/*PrintMe Prints ll*/
func PrintMe(head *ListNode) {
	current := head
	for current.Next != nil {

		fmt.Printf("%v->", current.Val)

		current = current.Next
	}
	fmt.Printf("%v\n", current.Val)
}

func HasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {

		if fast == nil {
			return false
		}

		if fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func DetectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next

	for slow != fast {

		if fast == nil {
			return nil
		}

		if fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	slow = slow.Next

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// func GetIntersectionNode(headA, headB *ListNode) *ListNode {

// }
