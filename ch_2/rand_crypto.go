/*
   隨機數分為以下幾類：
           1.偽隨機數 - 滿足第一個條件的隨機數。
           2.密碼學安全的偽隨機數 - 同時滿足前兩個條件的隨機數。可以通過密碼學安全偽隨機數生成器計算得出。
           3.真隨機數 -同時滿足三個條件的隨機數。
*/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 真隨機數
func main() {
	// 生成 10 個 [0, 100) 範圍的真隨機數。
	for i := 0; i < 10; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Println(result)
	}
}
