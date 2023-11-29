package main

import (
	"fmt"
	"golang_prac/note"
	"golang_prac/util"
)

var A = util.F("main.A")

func init() {
	util.F("main.init1")
}
func init() {
	util.F("main.init2")
}

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
	note.Defer()		// 3.6 defer
	note.DeferRecover()  // 3.6.2 defer, recover
	fmt.Println("note.DeferRecover() 之後還在運行")
	note.Array()		// 4.1 數組
	note.Slice()		// 4.2 切片
	note.Map()		// 4.3 map
	note.TypeDefinitionAndTypeAlias() // 4.4 類型定義與類型別名
	note.Struct()		// 4.5 結構體
}
