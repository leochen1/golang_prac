package note

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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


// 6.3 字符串常見操作 (strings包常見函數)
// strings.Contains(s, substr string) bool  // 判斷字符串 s 是否包含 substr
// strings.ContainsAny(s, chars string) bool  // 判斷字符串 s 是否包含 chars 中的任意字符
// strings.Count(s, sep string) int  // 統計字符串 s 中有幾個不重疊的 sep, 沒有則返回 -1
// strings.HasPrefix(s, prefix string) bool  // 判斷字符串 s 是否以 prefix 開頭
// strings.HasSuffix(s, suffix string) bool  // 判斷字符串 s 是否以 suffix 結尾
// strings.Index(s, sep string) int  // 返回字符串 s 中第一次出現 sep 的索引值，沒有則返回 -1
// strings.LastIndex(s, sep string) int  // 返回字符串 s 中最後一次出現 sep 的索引值，沒有則返回 -1
// strings.Replace(s, old, new string, n int) string  // 將字符串 s 中的前 n 個 old 替換為 new, n < 0 則替換所有
// strings.Split(s, sep string) []string  // 將字符串 s 按照 sep 分割為字符串數組
// strings.ToLower(s string) string  // 將字符串 s 中的所有字符轉換為小寫
// strings.ToUpper(s string) string  // 將字符串 s 中的所有字符轉換為大寫
// strings.TrimSpace(s string) string  // 將字符串 s 前後的空格去掉
// strings.Trim(s string, cutset string) string  // 將字符串 s 前後的 cutset 去掉
// strings.TrimLeft(s string, cutset string) string  // 將字符串 s 左邊的 cutset 去掉
// strings.TrimRight(s string, cutset string) string  // 將字符串 s 右邊的 cutset 去掉
// strings.Repeat(s string, count int) string  // 將字符串 s 重複 count 次
// strings.Join(str []string, sep string) string  // 將字符串數組 str 用 sep 連接
// strings.EqualFold("abc", "ABC")  // 不區分大小寫比較字符串
// strings.Fields(s string) []string  // 將字符串 s 按照空格分割為字符串數組
func PackageStrings() {
	fmt.Println(strings.Contains("advantech.com.tw", "advantech")) // true
	fmt.Println(strings.Index("advantech.com.tw", "c")) // 7
	fmt.Println(strings.Replace("advantech.com.tw", "advantech", "adv", -1)) // adv.com.tw
	fmt.Println(strings.Fields("ha ha \n ha\tha")) // [ha ha ha ha]
	fmt.Println(strings.Trim("#&^\nwww.www.www.&^#", "@#$%^&*")) // www.www.www.
}


// 6.4 中文字符常見操作 (utf8包常見函數)
// utf8.RuneCountInString(s string) int  // 返回字符串 s 中有幾個字符
// utf8.RuneLen(r rune) int  // 返回字符 r 的字節長度
// utf8.DecodeRuneInString(s string) (r rune, size int)  // 返回字符串 s 中第一個字符的值和字節長度
// utf8.DecodeLastRuneInString(s string) (r rune, size int)  // 返回字符串 s 中最後一個字符的值和字節長度
// utf8.EncodeRune(p []byte, r rune) int  // 將字符 r 寫入字節切片 p 中，返回寫入的字節長度
// utf8.RuneCount(p []byte) int  // 返回字節切片 p 中有幾個字符
// utf8.Valid(p []byte) bool  // 判斷字節切片 p 是否是合法的 utf8 編碼
// utf8.ValidString(s string) bool  // 判斷字符串 s 是否是合法的 utf8 編碼
// uth8.FullRune(p []byte) bool  // 判斷字節切片 p 中是否包含一個完整的字符
// utf8.FullRuneInString(s string) bool  // 判斷字符串 s 中是否包含一個完整的字符
func PackageUtf8() {
	fmt.Println(utf8.RuneCountInString("advantech.com.tw")) // 16
	fmt.Println(utf8.ValidString("advantech.com.tw")) // true
	fmt.Println(utf8.FullRuneInString("advantech.com.tw")) // true
}









