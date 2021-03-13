// 环形链表，判断是否有环，返回bool
//快慢指针
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}
	low := head
	fast := head
	for {
		if fast.next == ni || fast == nil { //防止越界
			return false
		}
		fast = fast.next.next
		low = low.next
		if fast == low {
			return true
		}
	}
}

// 哈希表
func isCycle(head *ListNode) bool {
	map:=map[&ListNode]struct{}{}  //?????????
	for head!=nil{
		if _,ok:=map[head];ok {
			return true
		}
		map[head]=struct{}{}
		head=head.next
	}
	return false
}

// 环形指针，判断是否有环，找出环的入口结点