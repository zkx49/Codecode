package main

// 求幂的方法
// 递归
func powerf1(x float64,n int)float64  {
	if n==0{
		return 1
	}else{
		return x*powerf1(x,n-1)
	}	
	
}
// 循环法
func powerf2(x float64,n int) float64 {
	ans:=1.0
	for n!=0{
		ans *=x
		n--
	}
	return ans
}
