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
	map:=map[*ListNode]struct{}{}  //初始化 结构体空值
	for head!=nil{
		if _,ok:=map[head];ok {
			return true
		}
		map[head]=struct{}{}
		head=head.next
	}
	return false
}
func isCycle1(head *ListNode)bool{
	hash:=make(map[*ListNode]bool)
	for head!=nil{
		if hash[head] {
			fmt.Println("cycle detect")
			break
		}
		hash[head]=true
	}
	return false
}

// 环形指针，判断是否有环，找出环的入口结点
// 哈希表
func detectCycle(head *Node)*Node{
	hashtable:=map[*Node]struct{}{}
	for head!=nil{
		if _,ok:=range hashtable;ok {
			return Node
		}
		hashtable[head]=struct{}{}
		head=head.next
	}
	return nil
}
// 快慢指针  把握思路，再理清逻辑，什么时候截止
func hasCycle(head *ListNode) *ListNode {
	slow,fast:=head,head
	for fast != nil{
		if fast.next==nil {
			return nil
		}
		fast=fast.next.next
		slow=slow.next
		if fast==slow {
			newhead:=head
			for newhead != slow{
				newhead=newhead.next
				slow=slow.next
			}
			return newhead
		}
	}
	return nil
}