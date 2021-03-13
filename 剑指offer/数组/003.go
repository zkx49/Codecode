// 连续数组的最大和，有正有负
// 思路：如果和大于零，则继续相加，否则从遍历到的数重新开始

func maxSubArray(nums []int) int {
	res := nums[0] //结果从第一个值开始
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		// 结果值和和值比较，返回最大值赋给结果值
		if res < sum {
			res = sum
		}
	}
	return res
}