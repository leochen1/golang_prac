package note

import (
	"fmt"
	"math/rand"
	"time"
)

// 6.1 隨機數
// Unix 時間 : UTC(世界標準時間)(1970-01-01 00:00:00) 到現在的秒數
// 納秒 : 1秒 = 10^9 納秒
// Go 獲取 Unix 納秒 : time.Now().UnixNano()
func RandNum() {
	seed := time.Now().UnixNano()  // 獲取當前時間的納秒
	r := rand.New(rand.NewSource(seed))  // 以當前時間的納秒作為種子數

	// 生成隨機數
	num := r.Intn(100) + 1 // 1~100
	fmt.Println("num = ", num)
}







