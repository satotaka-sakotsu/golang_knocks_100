/*
	命題
	"Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，
	取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
*/

package main

import (
	"strings"
	"fmt"
)

func main(){
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	str = strings.Replace(str, ".", "", -1)
	wordSlice := strings.Split(str, " ")

	wordIndexMap := make(map[string]int)
	for i, v := range wordSlice {
		switch i + 1 {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			wordIndexMap[string(v[0])] = i
		default:
			wordIndexMap[string(v[0:2])] = i
		}
	}

	fmt.Println(wordIndexMap)
	// 	map[Al:12 Ar:17 B:4 Be:3 C:5 Ca:19 Cl:16 F:8 H:0 He:1 K:18 Li:2 Mi:11 N:6 Na:10 Ne:9 O:7 P:14 S:15 Si:13]
}
