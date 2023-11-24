package main

import (
	"fmt"
	"golang_prac/note"
)

func main() {
	note.EscapedCharacters()     // 2.1 轉義字符
	note.VariablesAndConstants() // 2.2 變量與常量
	fmt.Println("note version : ", note.Version)
	note.BasicDataTypes()  // 2.3 基本數據類型
	note.Pointer()		 // 2.4 指針
	note.FmtVerbs()		 // 2.5 fmt占位符
	note.Operator()		 // 2.6 運算符
	note.IfElse()  	// 3.1 if else
	note.SwitchCase()		// 3.2 switch
	note.For()		// 3.3 for
	note.LabelAndGoto()	 // 3.4 label 與 goto
	note.Function()		// 3.5 函數, 匿名函數
	note.Function2()  // 3.5.2 函數, 直接調用函數
}
