/*
	命題
	"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
*/

package main

import (
	"strings"
	"fmt"
)

func main(){
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, ".", "", -1)
	wordSlice := strings.Split(str, " ")

	var wordCountSlice []int
	for _, v := range wordSlice {
		wordCountSlice = append(wordCountSlice, len(v))
	}

	fmt.Println(wordCountSlice)
	// => [3 1 4 1 5 9 2 6 5 3 5 8 9 7 9]
}
