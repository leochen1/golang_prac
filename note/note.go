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

	fmt.Println("4. 回車") // 一般與換行一起使用 \r\n
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
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7)
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v, v7=%v \n", v1, v2, v3, v4, v5, v6, v7)

	fmt.Println("2. 常量聲明")
	const (
		c1 = 8
		c2 = iota // 當前行數, iota是一個特殊的常量，可以被編譯器修改的常量, 從0開始，每增加一行加1, 直到下一個const出現
		c3 = iota
		c4 // 省略值時，默認與前一個值相同
		c5 = 12
		c6 // 省略值時，默認與前一個值相同
	)
	fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v \n", c1, c2, c3, c4, c5, c6)

	v = 200
	println("v = ", v)
}

// 2.3 基本數據類型
func BasicDataTypes() {
	fmt.Println("2.3.1 整數型")
	var (
		n1      = 5
		n2 int8 = 127
		n3 uint16
		n4 = 0b0101 // 二進制
		n5 = 0o777  // 八進制
		n6 = 0xAF   // 十六進制
	)
	fmt.Printf("n1=%v, type=%T \n", n1, n1)
	fmt.Printf("n2=%v, type=%T \n", n2, n2)
	fmt.Printf("n3=%v, type=%T \n", n3, n3)
	fmt.Printf("n4=%v, type=%T \n", n4, n4)
	fmt.Printf("n5=%v, type=%T \n", n5, n5)
	fmt.Printf("n6=%v, type=%T \n", n6, n6)

	fmt.Println("2.3.2 浮點型")
	var (
		f1         = 1.0
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
	c6 := 'A' - 'a' // 這裡的減法是將字符的碼值相減
	c7 := 'x'
	c8 := c7 + c6 // 這裡的加法是將字符的碼值相加, 但是結果是int類型, 並不是字符類型, 所以要用%c輸出, 才能看到字符,
	fmt.Printf("c8的碼值=%v, 這個碼值對應的字符是=%c, type=%T \n", c8, c8, c8)

	fmt.Println("2.3.5 布林值")
	var b1 bool
	fmt.Printf("b1=%v, type=%T \n", b1, b1)

	fmt.Println("2.3.6 字符串")
	var s1 = "Hello"
	println(s1 + "world")
	println(s1, "world")
	fmt.Println("len(s1) =", len(s1))

	s := `
		var (
		c1 byte
		c2 rune
		c3 = '0'
		c4 = '中'
		c5 = 22345
	)`
	fmt.Println(s) // 這裡的換行符會被保留
}

// 取址符 : &(獲取當前變量的內存地址)
// 取值符 : *(訪問地址指向的值)
// 數具類型 : *指向的類型
func increase(n int) {
	n++                                                         // 這裡的n是一個局部變量, 會在函數調用完畢後被回收, n = n + 1
	fmt.Printf("\nincrease結束時 n = %v \n n 的內存地址為 %v \n", n, &n) //&n 是 n 的內存地址
}

func increase2(n *int) {
	*n++                                                                           // 這裡的n是一個指針, 會將n的內存地址指向的值 + 1
	fmt.Printf("\nincrease2結束時 n = %v \n n 的內存地址為 %v \n n 指向的值為 %v \n", n, &n, *n) //&n 是 n 的內存地址, *n 是 n 的內存地址指向的值
}

// 2.4 指針
func Pointer() {
	var src = 2022
	increase(src)

	increase2(&src)                                                            // &src 是 src 的內存地址
	fmt.Printf("調用 increase2(src)之後, src = %v \n src 的內存地址為 %v \n", src, &src) //&src 是 src 的內存地址

	var ptr = new(int)                                                                            // new(T) 會為類型T分配一片內存空間, 並返回一個指向這片內存空間的指針
	fmt.Printf("調用 increase2(ptr)之後, ptr = %v \n ptr 的內存地址為 %v \n ptr 指向的值為 %v", ptr, &ptr, *ptr) //&ptr 是 ptr 的內存地址, *ptr 是 ptr 的內存地址指向的值
}

// 2.5 fmt 格式字符
func FmtVerbs() {
	fmt.Println("\n2.5.1 通用")
	fmt.Println("%%")

	fmt.Println("\n2.5.2 整數")
	i := 123
	fmt.Printf("%U\n", i) // 16進制
	fmt.Printf("%b\n", i) // 2進制
	fmt.Printf("%o\n", i) // 8進制
	fmt.Printf("%d\n", i) // 10進制
	fmt.Printf("%x\n", i) // 16進制
	fmt.Printf("%X\n", i) // 16進制
	fmt.Printf("%c\n", i) // Unicode碼值對應的字符
	fmt.Printf("%q\n", i) // Unicode碼值對應的字符, 用單引號包裹

	fmt.Println("\n2.5.3 浮點數")
	f := 123.456
	fmt.Printf("%f\n", f)     // 10進制
	fmt.Printf("%.2f\n", f)   // 10進制, 保留2位小數
	fmt.Printf("%20f\n", f)   // 10進制, 佔20位, 不足的部分用空格補齊
	fmt.Printf("%20.2f\n", f) // 10進制, 佔20位, 保留2位小數, 不足的部分用空格補齊
	fmt.Printf("%b\n", f)     // 指數為2的冪次方的科學計數法
	fmt.Printf("%e\n", f)     // 指數為10的冪次方的科學計數法
	fmt.Printf("%E\n", f)     // 指數為10的冪次方的科學計數法
	fmt.Printf("%g\n", f)     // 用%e或%f中較短的輸出表示, 並且不會有精確度的損失
	fmt.Printf("%G\n", f)     // 用%E或%f中較短的輸出表示, 並且不會有精確度的損失

	fmt.Println("\n2.5.4 布林值")
	fmt.Printf("%t\n", f == 123.456) // true

	fmt.Println("\n2.5.5 字符串或 byte 切片")
	s := "Hello World"
	fmt.Printf("%s\n", s) // 字符串
	fmt.Printf("%q\n", s) // 字符串, 用雙引號包裹
	fmt.Printf("%x\n", s) // 16進制
	fmt.Printf("%X\n", s) // 16進制

	fmt.Println("\n2.5.6 指針")
	p := &s
	fmt.Printf("%p\n", p) // 16進制, 指針的內存地址
}

// 2.6 運算符
func Operator() {
	fmt.Println("\n2.6.1 算術運算符")
	fmt.Println("8%3 =", 8%3)

	fmt.Println("\n2.6.2 位元運算符")
	var b uint8 = 0b00111100
	fmt.Printf("b>>2 = %b\n", b>>2) // 右移2位
	fmt.Printf("b<<2 = %b\n", b<<2) // 左移2位

	var b1 = 0b00111100
	var b2 = 0b00001101
	fmt.Printf("b1&b2 = %b\n", b1&b2) // 位元與
	fmt.Printf("b1|b2 = %b\n", b1|b2) // 位元或
	fmt.Printf("b1^b2 = %b\n", b1^b2) // 位元異或

	fmt.Println("\n2.6.3 邏輯運算符")
	var b3 = true
	var b4 = false
	fmt.Printf("b3&&b4 = %v\n", b3 && b4) // 邏輯與
	fmt.Printf("b3||b4 = %v\n", b3 || b4) // 邏輯或
	fmt.Printf("!b3 = %v\n", !b3)         // 邏輯非
}

// 3.1 if else
func IfElse() {
	var age uint
	fmt.Println("請輸入年齡: ")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("少兒不宜")
	} else if age < 18 {
		fmt.Println("青少年不宜")
	} else {
		fmt.Println("成人不宜")
	}
}

// 3.2 switch case
func SwitchCase() {
	var age uint
	fmt.Println("請輸入年齡: ")
	fmt.Scanln(&age)
	switch {
	case age < 13:
		fmt.Println("少兒不宜")
		fallthrough // 穿透, 繼續執行下一個case
	case age < 18:
		fmt.Println("青少年不宜")
	default:
		fmt.Println("成人不宜")
	}

	var weekday uint8
	fmt.Println("請輸入星期幾: ")
	fmt.Scanln(&weekday)
	switch weekday {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	default:
		fmt.Println("其他天")
	}
}

// 3.3 for 循環
func For() {
	fmt.Println("\n3.3.1 無限循環")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 10 {
			fmt.Println()
			break
		}
	}
	fmt.Println("\n3.3.2 條件循環")
	i = 1
	for i < 10 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println("\n3.3.2 標準循環")
	for j := 1; j < 11; j++ {
		fmt.Print(j, "\t")
	}
}

// 3.4 label 與 goto
func LabelAndGoto() {
	// outside:
	fmt.Println("\n3.4.1 label")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				// break outside
			}
		}
		fmt.Println()
	}
	fmt.Println("\n3.4.2 goto")
	fmt.Print("1")
	if i := 1; i == 1 {
		goto four // 跳轉到four標籤, 不推薦使用
	}
	fmt.Print("2")
	fmt.Print("3")
four:
	fmt.Print("4")
	fmt.Print("5")
}

// 3.5 函數
func Function() {
	// 匿名函數
	var getRes = func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}
	res1, res2 := getRes(2, 3)
	fmt.Println("res1 =", res1, "res2 =", res2)
	// fmt.Printf("getRes=%v, Type of getRes=%T\n", getRes, getRes)
}

func Function2() {
	// 直接調用函數
	res1, res2 := func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	fmt.Println("res1 =", res1, "res2 =", res2)
}

// 3.6 defer
// defer 會將函數推遲到當前函數執行完畢後再執行
// defer 會將函數壓入一個棧中, 當函數執行完畢後, 會從棧中取出函數, 並執行
// defer 會 "先進後出", 類似於棧
// defer 會先執行參數, 再執行函數
// defer 會先執行後定義的, 再執行先定義的
// defer 會先執行外層函數, 再執行內層函數
// defer 會先執行return, 再執行defer
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次調用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函數被第%v次調用\n", i)
		return i
	}
}
func Defer() int {
	f := deferUtil()
	//defer f(f(3))  // 這裡的f(f(3))會先執行, 但是不會立即執行, 會被壓入棧中, 等到函數執行完畢後, 再從棧中取出, 並執行
	defer f(1) // 這裡的f(1)會先執行, 但是不會立即執行, 會被壓入棧中, 等到函數執行完畢後, 再從棧中取出, 並執行
	defer f(2)
	defer f(3)
	return f(4)
}
func DeferRecover() {
	defer func() {
		err := recover() // recover()必須在defer中執行, 並且必須在異常發生後執行
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}



