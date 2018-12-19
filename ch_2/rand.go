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
)

// 偽隨機數
// 隨機種子( Seed, 開始種子 ) 默認初始種子總是從 1 開始。
func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
