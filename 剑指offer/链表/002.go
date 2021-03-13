// 输入一个链表，按链表从尾到头的顺序返回一个ArrayList
type ListNode struct {
	value int
	next  *ListNode
}

func fmtListNode(head *ListNode) {
	l := list.New()
	for head != nil {
		l.PushFront(head.value)
		head = head.next
	}
	for item := l.Front(); item != nil; item = item.Next {
		fmt.Println(item.value)
	}
}