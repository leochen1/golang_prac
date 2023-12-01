package note

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 6.1 隨機數
// Unix 時間 : UTC(世界標準時間)(1970-01-01 00:00:00) 到現在的秒數
// 納秒 : 1秒 = 10^9 納秒
// Go 獲取 Unix 納秒 : time.Now().UnixNano()
func RandNum() {
	seed := time.Now().UnixNano()       // 獲取當前時間的納秒
	r := rand.New(rand.NewSource(seed)) // 以當前時間的納秒作為種子數

	// 生成隨機數
	num := r.Intn(100) + 1 // 1~100
	fmt.Println("num = ", num)
}

// 6.2 字符串類型轉換
// 6.2.1 fmt 包
// fmt.Sprintf(fmtstr, ...)  // 其他類型轉換為字符串, 用法同 fmt.Printf()
// fmt.Sscanf(str, fmtstr, ...)  // 字符串解析為其他類型，用法同 fmt.Scanf()
// 6.2.2 strconv 包
// strconv.Itoa(i int) string  // 整數轉換為字符串
// strconv.Atoi(s string) (i int, err error)  // 字符串轉換為整數
// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string  // 浮點數轉換為字符串
// strconv.ParseFloat(s string, bitSize int) (f float64, err error)  // 字符串轉換為浮點數
// strconv.FormatBool(b bool) string  // 布爾值轉換為字符串
// strconv.ParseBool(str string) (value bool, err error)  // 字符串轉換為布爾值
// strconv.FormatInt(i int64, base int) string  // 整數轉換為字符串
// strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)  // 字符串轉換為整數
// strconv.FormatUint(i uint64, base int) string  // 整數轉換為字符串
// strconv.ParseUint(s string, base int, bitSize int) (i uint64, err error)  // 字符串轉換為整數

func StrConv() {
	i1 := 123
	s1 := "advantech.com.tw"
	s2 := fmt.Sprintf("%d@%s", i1, s1) // 整數轉換為字符串
	fmt.Println("s2 = ", s2)

	var (
		i2 int
		s3 string
	)
	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3) // 字符串解析為其他類型
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功解析了%d個數據\n", n)
	fmt.Println("i2 = ", i2)
	fmt.Println("s3 = ", s3)

	s4 := strconv.FormatInt(123, 4)  // 整數轉換為字符串
	fmt.Println("s4 = ", s4)
	u1, err := strconv.ParseUint(s4, 4, 0)  // 字符串轉換為整數
	if err != nil {
		panic(err)
	}
	fmt.Println("u1 = ", u1)
}






