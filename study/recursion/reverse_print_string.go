package recursion

import "fmt"

/*
反向打印字符串
*/
func reversePrintString(s string) {
	if s == "" {
		return
	}
	reversePrintString(s[1:])
	fmt.Println(string(s[0]))
}
