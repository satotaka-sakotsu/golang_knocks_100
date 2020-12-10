/*
	命題
	「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ。
*/

package main

import (
	"strings"
	"fmt"
)

func main(){
	str := "パタトクカシーー"
	slice := strings.Split(str, "")

	var oddStr string
	for i := range slice {
		if (i + 1) % 2 == 0 {
			continue
		}
		oddStr += slice[i]
	}

	fmt.Println(oddStr)
	// => パトカー
}
