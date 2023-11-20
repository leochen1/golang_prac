package main

import (
	"fmt"
	"golang_prac/note"
)

func main() {
	note.EscapedCharacters()     // 2.1 轉義字符
	note.VariablesAndConstants() // 2.2 變量與常量
	fmt.Println("note version : ", note.Version)
}
