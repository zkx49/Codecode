// 一、理解指针或引用的含义
// 无论指针还是引用，都是存储所指对象的内存地址。
// 将某个变量赋值给指针，实际上就是将这个变量的地址赋值给指针。
// 二、警惕指针丢失或内存泄漏
// 插入结点时，一定要注意操作顺序，要将结点x的next指针指向结点b，再把结点a的指针指向结点x  X.next-->B..
// 删除结点时，也一定要记得手动释放内存空间。
// 三、利用哨兵简化实现难度
// 四、重点留意边界条件处理
// 例：如果链表为空，代码是否能正常工作？
//        如果链表只包含一个结点，代码是否能正常工作？
//        如果链表只包含两个结点，代码是否能正常工作？
//        代码逻辑在处理头结点和尾结点的时候，是否能正常工作？
// 五、举例画图，辅助思考
// 六、无捷径，多写多练
type ListNode struct {
	value interface{}
	next  *ListNode
}
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{]}) *ListNode {
	return &ListNode{v,nil}
}
func NewLinkedList() *LinkedList {
	return &LinkedList{NewLinkedList(0),0}
}