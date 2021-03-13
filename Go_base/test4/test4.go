func replacespace(s string) string {
	res := ""
	for _, v := range s {
		if v=='' {
			res+="%20"	
		}else{
			res+=string(v)
		}
	}
	return res
}
// 两种循环遍历string是不一样的
for i := 0; i < len(s); i++ {
	//utf-8 遍历 int8类型 byte类型
}
for _, v := range v {
	// unicode遍历 int32 rune类型
}