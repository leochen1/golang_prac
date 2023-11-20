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


// 2.3 基本數據類型
func BasicDataTypes() {
	fmt.Println("2.3.1 整數型")
	var (
		n1 = 5
		n2 int8 = 127
		n3 uint16
		n4 = 0b0101  // 二進制
		n5 = 0O777   // 八進制
		n6 = 0xAF	// 十六進制
	)
	fmt.Printf("n1=%v, type=%T \n", n1, n1)
	fmt.Printf("n2=%v, type=%T \n", n2, n2)
	fmt.Printf("n3=%v, type=%T \n", n3, n3)
	fmt.Printf("n4=%v, type=%T \n", n4, n4)
	fmt.Printf("n5=%v, type=%T \n", n5, n5)
	fmt.Printf("n6=%v, type=%T \n", n6, n6)

	fmt.Println("2.3.2 浮點型")
	var (
		f1 = 1.0
		f2 float32 = 1
		f3 float64 = 1
	)
	fmt.Printf("f1=%v, type=%T \n", f1, f1)
	fmt.Printf("f2=%v, type=%T \n", f2, f2)
	fmt.Printf("f3=%v, type=%T \n", f3, f3)

	fmt.Println("2.3.3 數值型數據類型轉換")
	n2 = int8(n3)
	fmt.Printf("n2=%v, type=%T \n", n2, n2)

	fmt.Println("2.3.4 字符型")
	var (
		c1 byte
		c2 rune
		c3 = '0'
		c4 = '中'
		c5 = 22345
	)
	fmt.Printf("c1的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c1, c1, c1)
	fmt.Printf("c2的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c2, c2, c2)
	fmt.Printf("c3的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c3, c3, c3)
	fmt.Printf("c4的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c4, c4, c4)
	fmt.Printf("c5的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c5, c5, c5)
	c6 := 'A' - 'a'  // 這裡的減法是將字符的碼值相減
	c7 := 'x'  
	c8 := c7 + c6  // 這裡的加法是將字符的碼值相加, 但是結果是int類型, 並不是字符類型, 所以要用%c輸出, 才能看到字符,
	fmt.Printf("c8的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c8, c8, c8)

	fmt.Println("2.3.5 布林值")
	var b1 bool
	fmt.Printf("b1=%v, type=%T \n", b1, b1)

	fmt.Println("2.3.6 字符串")
	var s1 = "Hello"
	println(s1 + "world")
	println(s1, "world")
	fmt.Println("len(s1) =", len(s1))
}


