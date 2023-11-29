package note

import (
	"fmt"
	"golang_prac/util"
)

var v int = 100

var A = util.F("note.A")

func init() {
	util.F("note.init1")
}
func init() {
	util.F("note.init2")
}

// 2.1 轉義字符
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

// 2.2 變量與常量
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

// 2.4 指針
func Pointer() {
	var increase = func(n int) {
		n++                                                         // 這裡的n是一個局部變量, 會在函數調用完畢後被回收, n = n + 1
		fmt.Printf("\nincrease結束時 n = %v \n n 的內存地址為 %v \n", n, &n) //&n 是 n 的內存地址
	}
	var increase2 = func(n *int) {
		*n++                                                                           // 這裡的n是一個指針, 會將n的內存地址指向的值 + 1
		fmt.Printf("\nincrease2結束時 n = %v \n n 的內存地址為 %v \n n 指向的值為 %v \n", n, &n, *n) //&n 是 n 的內存地址, *n 是 n 的內存地址指向的值
	}
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

// 3.7 init
// init 函數會在 main 函數之前執行, 並且每個包都可以有多個 init 函數, 但是 init 函數不能被其他函數調用, 也不能被其他函數引用, 也不能有返回值, 也不能有參數
// 1. 每個包可以有自己的 init 函數, 並且可以有多個
// 2. 執行順序(取決於包的依賴關係) : 被依賴包的全局變量 > 被依賴包的 init 函數 > main包的全局變量 > main包的 init 函數 > main 函數

// 3.8 包
// 包名與目錄名可以不一致, 但是一般情況下, 包名與目錄名一致, 並且包名與目錄名都是小寫字母, 並且包名與目錄名不能有空格, 不能有特殊字符, 不能有數字開頭, 不能與內置關鍵字相同
// 1. GO程式是由包構成的, 程式從 main 包開始運行
// 2. 引入包的語法 : import "包的路徑"
// 3. 包的別名 : import 別名 "包的路徑"

// 4.1 數組
// 1. 數組是值類型, 數組的長度是數組類型的一部分, 也就是說 [5]int 和 [10]int 是不同的類型
// 2. 數組的長度必須是常量, 並且長度是數組類型的一部分, 所以數組的長度是不能改變的
// 3. 數組的元素可以是任意類型, 包括值類型和引用類型, 但是不能混用
// 4. 數組可以通過下標訪問, 下標從0開始, 到長度減1
func Array() {
	// 4.1.1 聲明
	var a = [...]int{1, 456, 789}
	a[0] = 123

	// 4.1.2 遍歷
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}

	// 4.1.3 for range
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}

	// 4.1.4 多維數組
	var twoDimensionalArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7},
		{8, 9, 10, 11},
	}
	for i, v := range twoDimensionalArray {
		for j, v2 := range v {
			fmt.Printf("twoDimensionalArray[%v][%v]=%v\n", i, j, v2)
		}
	}
}

// 4.2 切片
// 切片是對數組的引用，切片本身並不存儲任何數據，他只是描述了底層數組中的一段
// 1. 切片的長度是可以變化的，因此切片是一種動態數組
// 2. 切片是一種引用類型，他的內部結構包含了三個元素：底層數組的指針、切片的長度（len）和切片的容量（cap）
// 3. 切片一般用於快速地操作一塊數據集合
// 4.2.1 聲明
// 1. 引用數組的一段
// 2. 引用切片的一段*(將指向同一個底層數組)
// 3. 分配內存空間make([]T, len, cap)
// 4.2.2 長度與容量
// len(s) : 切片的長度
// cap(s) : 切片的容量
func Slice() {
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4] // 這裡的1是起始下標, 4是結束下標, 但是不包含結束下標對應的元素, [0:len(array)] = [:]
	s1[0] = 0
	fmt.Println("array = ", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array = ", array)
	var s3 []int // nil 切片, 長度和容量都是0, 並且沒有底層數組, 一般用於判斷切片是否初始化, 不能對nil切片進行操作, 但是可以對nil切片賦值
	fmt.Println("does s3==nil? ", s3 == nil)
	s3 = make([]int, 3, 5) // make([]T, len, cap) 這裡的len是切片的長度, cap是切片的容量, 並且len <= cap, 如果不寫 cap, 那麼 cap = len
	fmt.Printf("len(s3) = %v, cap(s3) = %v", len(s3), cap(s3))
	s4 := []int{1, 2, 3} // 這裡的[]int是切片類型, {1,2,3}是切片的值, 並且這裡的len = cap = 3
	fmt.Printf("s4 = %v, len(s4) = %v, cap(s4) = %v", s4, len(s4), cap(s4))

	// 4.2.3 追加元素
	s1 = append(s1, 6, 7, 8) // 底層創建了新的數組，不再引用原數組，append(s, x...) 這裡的s是切片, x是元素, 這裡的x可以是多個元素, 並且可以是切片, 但是切片需要展開
	s1[4] = 0
	fmt.Println("array = ", array)
	fmt.Println("s1 = ", s1)

	s5 := append(s1, s2...) // Fix: Use '...' to expand the elements of s2 , 表示將 s2 切片中的所有元素作為參數傳遞。
	fmt.Println("s5 = ", s5)

	// 4.2.4 拷貝切片
	s6 := []int{1, 1, 1}
	copy(s6, s5) // copy(dst, src) 這裡的dst是目標切片, src是源切片, 並且會將src中的元素拷貝到dst中, 並且會覆蓋dst中的元素, 並且會以len(src)和len(dst)中的最小值作為拷貝的長度
	fmt.Println("s6 = ", s6)
	copy(s5, s6)  // 這裡的len(src) = 3, len(dst) = 7, 所以只會拷貝3個元素, 並且會覆蓋s5中的元素
	fmt.Println("s5 = ", s5)  

	// 4.2.5 string 與 []byte
	// string 是不可變的, 也就是說不能通過 s[0] = 'a' 這種方式來修改字符串中的字符, 但是可以通過 []byte 來修改, 並且 []byte 可以和 string 互相轉換
	s7 := "Hello 世界"
	s8 := []byte(s7) // 將字符串轉換為 []byte
	fmt.Printf("s8 = %s", s8)
	fmt.Printf("s8 = %v", s8)
	s8[0] = 'h'
	fmt.Printf("s8 = %s", s8)

	for i, v := range s7 {
		fmt.Printf("s7[%d] = %c\n", i, v)
	}

	key:= util.SelectByKey("註冊", "登錄", "退出")
	fmt.Println("key =", key)
}


// 4.3 map
// 1. map 是一種無序的 key-value 對集合, key是唯一的, 並且key是可以比較的, 並且key只能是值類型, 並且value可以是任意類型
// 2. map 是引用類型, 並且可以使用 make() 來創建, 並且可以使用 len() 來獲取元素個數
// 3. map 的容量達到後, 再增加元素, map 的容量會自動擴容
// 4. map 的value為引用類型時, 修改了原始值, map 中的值也會跟著改變
func Map() {
	var m1 map[string]string
	fmt.Println("m1 = nil?", m1==nil)

	m1 = make(map[string]string)  // make(map[string]string, 2) 這裡的2是容量, 並不是長度, 並且容量是可選的, 如果不寫 (默認1), 那麼容量 = 長度
	m1["name"] = "Tom"
	m1["age"] = "18"
	m1["height"] = "180"
	fmt.Println("m1 =", m1)

	m2 := map[string]string{
		"name": "Tom",
		"age": "18",
	}
	fmt.Println("m2 =", m2)

	v, ok := m2["name"]  // 這裡的ok是一個bool值, 如果key存在, 那麼ok = true, 如果key不存在, 那麼ok = false
	if ok {
		fmt.Println("m2[\"name\"] =", v)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m2, "age")  // 刪除key
	fmt.Println("m2 =", m2)

	for key, value := range m1 {
		fmt.Printf("m1[%v] = %v\n", key, value)
	}

	m1 = nil  // 將map賦值為nil, 並不會釋放map佔用的內存, 但是map中的元素會被回收
	m2 = make(map[string]string)  // 這裡的m2會重新分配內存, 並且map中的元素會被回收
	fmt.Println("m1 =", m1)
	fmt.Println("m2 =", m2)
}


// 4.4 自定義數據類型&類型別名
// 1. type 類型名 類型
// 2. 類型別名只會在程序中存在, 在編譯後不會有類型別名的存在
// 3. 類型別名和類型在使用上完全一致, 並且可以相互轉換, 但是類型別名和類型在使用上是不同的, 不能相互轉換
func TypeDefinitionAndTypeAlias() {
	fmt.Println("\n4.4.1 自定義數據類型")
	type mesType uint16  // 自定義數據類型
	var u1000 uint16 = 1000  
	var textMes mesType = mesType(u1000)  // 需做類型轉換
	fmt.Printf("textMes = %v, type = %T\n", textMes, textMes)

	fmt.Println("\n4.4.2 類型別名")
	type myUnit16 = uint16  // 類型別名只會在程序中存在, 在編譯後不會有類型別名的存在
	var myu16 myUnit16 = 1000  // 不需做類型轉換
	fmt.Printf("myu16 = %v, type = %T\n", myu16, myu16)
}


// 4.5 結構體
// 1. 結構體是"值""類型, 通過 . 訪問
// 2. 結構體的每個字段上，可以寫一個tag，該tag可以通過反射機制獲取，常用於序列化和反序列化
// 3. 結構體中的字段名必須唯一
// 結構體指針, 注意 "." 優先級高於 "&" / "*"
// 使用時可以簡寫(隱式間接引用) : p1.name = "Tom" 等價於 (*p1).name = "Tom"
//  可以使用 "&" 前綴快速聲明結構體指針 : p2 := &person{} 
type User struct {
	Name string
	Id uint32
}
type Account struct {
	User  // 匿名字段, 這裡的User是類型, 並不是字段名, 這裡的User類型是User結構體, 並且User結構體中的字段會被提升到Account結構體中
	password string
}
type Contact struct {
	*User  // 匿名字段, 這裡的User是類型, 並不是字段名, 這裡的User類型是*User結構體指針, 並且User結構體中的字段會被提升到Contact結構體中
	Remark string
}
func Struct() {
	var u1 User = User {
		Name: "張三",
	}
	u1.Id = 10000
	var u2 *User = &User {
		Name: "李四",
	}
	u2.Id = 10001  // (*u2).Id = 10001

	var a1 = Account {
		User: User {
			Name: u1.Name,
		},
		password: "123456",
	}
	var c1 *Contact = &Contact {
		User: &User{
			Id: u2.Id,
		},  
		Remark: "備註",
	}
	c1.Name = "趙六"  //c1.User.Name = "趙六"  //沒有重複字段時, 可以簡寫
	fmt.Println("a1 = ", a1)
	fmt.Println("c1 = ", c1)
	fmt.Println("c1.User = ", c1.User)
}









