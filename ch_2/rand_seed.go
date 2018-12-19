/*
   隨機數分為以下幾類：
           1.偽隨機數 - 滿足第一個條件的隨機數。
           2.密碼學安全的偽隨機數 - 同時滿足前兩個條件的隨機數。可以通過密碼學安全偽隨機數生成器計算得出。
           3.真隨機數 -同時滿足三個條件的隨機數。
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 偽隨機數，採用當前時間作為開始種子。
func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(int64(time.Now().UnixNano()))

	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
}
