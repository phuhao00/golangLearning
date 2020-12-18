package likou

//实现一个算法，确定一个字符串 s 的所有字符是否全都不同
func isUnique(astr string) bool {
	arrTmp:=[]rune(astr)
	checkMap:=make(map[rune]bool)
	for _, val := range arrTmp {
		if checkMap[val]{
			return false
		}
		checkMap[val]=true
	}
	return true
}