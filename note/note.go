package note

import "fmt"

var v int = 100

//2.1 轉義字符
func EscapedCharacters() {
	fmt.Println("1. 雙引號")
	fmt.Println("\"Hello\"")

	fmt.Println("2. 反斜線")
	fmt.Println("\\Hello\\")

	fmt.Println("3. 換行")
	fmt.Println("Hello\nWorld")

	fmt.Println("4. 回車")  // 一般與換行一起使用 \r\n
	fmt.Println("Hello\rWorld")

	fmt.Println("5. 水平制表符")
	fmt.Println("Hello\tWorld")

	fmt.Println("6. 垂直制表符")
	fmt.Println("Hello\vWorld")

	fmt.Println("7. 警報鈴聲")
	fmt.Println("\aHello\aWorld\a")

	fmt.Println("8. 退格")
	fmt.Println("Hello\bWorld")

	fmt.Println("9. 16進制")
	fmt.Println("\x41\x42\x43")

	fmt.Println("10. 8進制")
	fmt.Println("\101\102\103")

	fmt.Println("11. Unicode")
	fmt.Println("\u4e16\u754c")
}

const (
	Version int = 100
)

//2.2 變量與常量
func VariablesAndConstants() {
	fmt.Println("1. 變量聲明")
	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4
	var (
		v5 = 5
		v6 int = 6
		v7 int
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7)
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v, v7=%v \n", v1, v2, v3, v4, v5, v6, v7)

	fmt.Println("2. 常量聲明")
	const(
		c1 = 8
		c2 = iota  // 當前行數, iota是一個特殊的常量，可以被編譯器修改的常量, 從0開始，每增加一行加1, 直到下一個const出現
		c3 = iota
		c4  // 省略值時，默認與前一個值相同
		c5 = 12
		c6  // 省略值時，默認與前一個值相同
	)
	fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v \n", c1, c2, c3, c4, c5, c6)

	v = 200
	println("v = ", v)
}



