// 字符串中的第一个唯一字符
// 方法一：使用哈希表存储频数  问题在于不能理解具有大小的空数组
// 空数组是要有确定的大小，这会浪费一定的内存
// 思想：遍历两次字符串，一次用数组保存其值的次数，然后再在字符串中遍历
func firstUniqChars(s string) int {
	cnt:=[26]int{}   //大小为26的整型空数据
	for _, ch := range s{ //ch为value
		cnt[ch-'a']++    //ch-'a'为数组下标，每次遍历到则增加其对应的值
	}
	for i,ch  := range s {
		if cnt[ch-'a']==1 {
			return i
		}
	}
	return -1
}
