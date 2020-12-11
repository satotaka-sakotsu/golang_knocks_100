/*
	命題
	与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
	* 英小文字ならば(219 - 文字コード)の文字に置換
	* その他の文字はそのまま出力
	この関数を用い，英語のメッセージを暗号化・復号化せよ．
*/

package main

import (
	"fmt"
)

func cipher(str string) string {
	var encodedStr string
	for _, v := range str {
		if v > 'a' && v < 'z' {
			encodedStr += string(219 - v)
		} else {
			encodedStr += string(v)
		}
	}

	return encodedStr
}

func main(){
	str := "foo bar baz. FOO BAR BAZ :)"
	encode := cipher(str)
	decode := cipher(encode)
	fmt.Println("元の文字", str)
	fmt.Println("暗号化", encode)
	fmt.Println("複合化", decode)
}
