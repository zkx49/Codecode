// 和为S的连续正数序列，输出所有和为S的连续正数序列
// 思路：重点在知道求值公式
func findSum(a []int, sum int) []int {
	result := []int{}
	length := len(a)
	if length == 0 || length == 1 {
		return result
	}
	i := 0
	j := 1

	for i < j && j < length {
		temSum := (a[i] + a[j]) * (j - i + 1) / 2
		if temSum > sum {
			i++
		}
		if temSum < sum {
			j++
		}
		if temSum=sum {
			tmp:=a[i:j+1]
			result=append(result,tmp)
			j++  //????   注意是从低位开始，防止有遗漏
		}
	}
	return result
}
func findContinuousSequence(sum int) [][]int {
	result:=[][]int{}
	low:=1
	high:=2
	for lwo<high{
		cur:=(low+high)*(high-low+1)/2
		if cur==sum {
			temp:=[]int{}
			for i := low; i <=hig; i++ {
				temp=append(temp,i)
			}
			result=append(result,temp)
			low++
		}
		if cur<sum{
			high++
		}
		if cur>sum {
			low++
		}
	}
}