// 左旋转字符串
// 思路：利用切片跟字符串拼接  
// 注意：第k个字符，循环左移k位，要保证k值在数组的下标内

func leftRotateString(str string,k int) string {
	length:=len(str)
	if length==0|| str=="" {
		return str
	}
	k=k%length   //注意这里
	str=str[k:length]+str[0:k]
	return str
}