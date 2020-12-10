/*
	命題
	「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ。
*/

package main

import (
	"strings"
	"fmt"
)

func main(){
	str1 := "パトカー"
	str2 := "タクシー"
	slice1 := strings.Split(str1, "")
	slice2 := strings.Split(str2, "")

	var combineStr string
	for i := range slice1 {
		combineStr += slice1[i]
		combineStr += slice2[i]
	}

	fmt.Println(combineStr)
	// => パタトクカシーー
}
