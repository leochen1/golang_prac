package util

import "fmt"

var i = 0
var F = func(s string) int {
	fmt.Printf("本次被%s\n", s)
	i++
	fmt.Printf("匿名工具函數被第%v次調用\n", i)
	return i
}

func SelectByKey(text ...string) (key int) {
	for i, v := range text {
		fmt.Printf("%v : %v\n", i, v)
	}
	fmt.Println("請輸入你的選擇")
	fmt.Scanln(&key)
	return key
}
