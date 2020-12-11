/*
	命題
	与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．
*/

package main

import (
	"strings"
	"fmt"
)

type gramType int

const (
	charaGram gramType = iota
	wordGram
)

func printBiGram(str string, gt gramType)  {
	separator := ""
	if gt == wordGram {
		separator = " "
	}

	list := strings.Split(str, separator)

	for i := range list {
		if i == 0 {
			continue
		}

		grams := []string{list[i - 1], list[i]}
		gram := joinGram(grams, gt)
		fmt.Println(gram)
	}
}

func joinGram(strList []string, gt gramType) string {
	if gt == wordGram {
		return strings.Join(strList, " ")
	} else {
		return strings.Join(strList, "")
	}
}

func main(){
	str := "I am an NLPer"

	fmt.Println("----- 文字bi-gram")
	printBiGram(str, charaGram)

	fmt.Println("----- 単語bi-gram")
	printBiGram(str, wordGram)
/* =>
	----- 文字bi-gram
	I
	 a
	am
	m
	 a
	an
	n
	 N
	NL
	LP
	Pe
	er
	----- 単語bi-gram
	I am
	am an
	an NLPer
*/
}
