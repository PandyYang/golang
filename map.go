package main

// func main() {
// 	m := map[string]string{
// 		"1":"1",
// 		"2":"2",
// 		"3":"3",
// 		"4": "4",
// 	}
// 	fmt.Println(m)
// 	m2 := make(map[string] int)//空的map
// 	fmt.Println(m,m2)
// 	for k,v := range m{
// 		fmt.Println(k,v)
// 	}
// 	numberName := m["1"]
// 	fmt.Println(numberName)

// 	//numberName,ok := m["2"]
// 	if numberName,ok := m["2"];ok {
// 		fmt.Println(numberName)
// 	}else {
// 		fmt.Println("key does not exists")
// 	}
// }

//字符串中最长的无重复的子字符串的长度
func lengthConDup(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
