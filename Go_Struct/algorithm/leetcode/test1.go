// 寻找最长不含有重复字符的子串，放回它的长度
// 思路：遍历字符串，用map存储值和下标，相同的值会查找成功，然后更改最大长度下标的位置
func lengthOf(s string) int {
	hash := make(map[rune]int)
	start := 0     //开始的时候和出现重复值的下表
	maxLength := 0 //记录最大长度
	for k, v := range s {
		if lastIndex, ok := hash[v]; ok && lastIndex >= start {
			start = lastIndex + 1 //下标
		}
		if i-start+1 >= maxLength { //长度
			maxLength = i - start + 1
		}
		hash[v] = k
	}
	return maxLength

}

// make(type, len,cap)

// 有效括号，（）[] {}
// 利用栈，入栈，出栈
// 利用哈希表，预先初始化键值，不在哈希表中则入栈，然后栈顶的元素符合哈希表初始的值出栈
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	hash := map[byte]byte{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if hash[s[i]] > 0 { //哈希表中找到
			if len(stack) == 0 || stack[len(stack)-1] != hash[s[i]] { //头和尾的边界判定
				return false
			}
			stack = stack[:len(stack)-1] //出栈
		} else {
			stack = append(stack, s[i]) //入栈
		}
	}
	return len(stack) == 0
}