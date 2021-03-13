// 和为S的两个数字，递增排序的队列，如果有多对，输出两个数乘积最小的,小的先输出
// 思考：为什么只找到一对就停止？
func findNumbersWithSum(a []int, sum int) []int {
	result := []int{}
	length := len(a)
	if length==0 {
		return result
	}
	i := 0
	j := length - 1
	for i < j {
		if a[i]+a[j] = sum {
			result = append(result,a[i], a[j])
			break
		}
		if a[i]+a[j] > sum {
			j--
		}
		if a[i]+a[j] < sum {
			i++
		} 
	}
	return result
}