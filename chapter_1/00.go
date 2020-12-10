/*
	命題
	文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ。
*/

package main

import (
	"strings"
	"fmt"
)

func main(){
	str := "stressed"
	slice := strings.Split(str, "")

	var reverseStr string
	for i := range slice {
		reverseStr += slice[len(slice) - 1 - i]
	}

	fmt.Println(reverseStr)
	// => desserts
}
