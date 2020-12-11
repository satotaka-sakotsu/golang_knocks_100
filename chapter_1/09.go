/*
	命題
	スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．
	ただし，長さが４以下の単語は並び替えないこととする．適当な英語の文
	（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）
	を与え，その実行結果を確認せよ．
*/

package main

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
)

func convert(str string) string {
	strSlice := strings.Split(str, " ")
	var convertedSlice []string

	for _, v := range strSlice {
		if len(v) > 4 {
			first := string(v[0])
			last := string(v[len(v) - 1])
			middle := shuffle(v[1:len(v) - 1])
			converted := first + middle + last
			convertedSlice = append(convertedSlice, converted)
		} else {
			convertedSlice = append(convertedSlice, v)
		}
	}

	return strings.Join(convertedSlice, " ")
}

func shuffle(str string) string {
	rand.Seed(time.Now().UnixNano())
	r := []rune(str)

	for i := range str {
		idx := rand.Intn(len(str))
		r[idx], r[i] = r[i], r[idx]
	}

	return string(r)
}

func main(){
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	fmt.Println(convert(str))
}
