/*
	命題
	"paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．
	さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．
*/

package main

import (
	"fmt"
	"strings"
)

func biGram(str string) []string {
	list := strings.Split(str, "")
	var grams []string

	for i := range list {
		if i == 0 {
			continue
		}

		grams = append(grams, list[i - 1] + list[i])
	}

	return grams
}

func uniqueGram(biGram []string) []string {
	unique := []string{}
	for _, gram := range biGram {
		strUnique := strings.Join(unique, " ")
		if !strings.Contains(strUnique, gram) {
			unique = append(unique, gram)
		}
	}

	return unique
}

func union(biGrams ...[]string) []string {
	var unionGram []string

	for _, biGram := range biGrams {
		if len(unionGram) == 0 {
			unionGram = biGram
		}

		for _, gram := range biGram {
			strUnionGram := strings.Join(unionGram, " ")
			if !strings.Contains(strUnionGram, gram) {
				unionGram = append(unionGram, gram)
			}
		}
	}

	return unionGram
}

func intersection(biGrams ...[]string) []string {
	var inGram []string

	for i, biGram := range biGrams {
		if i == 0 {
			continue
		}
		prevGram := biGrams[i - 1]

		for _, gram := range biGram {
			strPrevGram := strings.Join(prevGram, " ")
			if strings.Contains(strPrevGram, gram) {
				inGram = append(inGram, gram)
			}
		}
	}

	return uniqueGram(inGram)
}

func difference(biGrams ...[]string) []string {
	var diffGram []string

	for i, currentBiGram := range biGrams {
		if i == 0 {
			continue
		}
		prevGram := biGrams[i - 1]

		for _, gram := range prevGram {
			strCurrentGram := strings.Join(currentBiGram, " ")
			if !strings.Contains(strCurrentGram, gram) {
				diffGram = append(diffGram, gram)
			}
		}
	}

	return uniqueGram(diffGram)
}

func isInclude(strSlice []string, keyword string) bool {
	str := strings.Join(strSlice, " ")
	return strings.Contains(str, keyword)
}

func main() {
	str1 := "paraparaparadise"
	str2 := "paragraph"

	biGramX := biGram(str1)
	biGramY := biGram(str2)

	uniqueGramX := uniqueGram(biGramX)
	uniqueGramY := uniqueGram(biGramY)

	fmt.Println("X bi-gram", biGramX)
	fmt.Println("Y bi-gram", biGramY)
	fmt.Println("和集合", union(uniqueGramX, uniqueGramY))
	fmt.Println("積集合", intersection(uniqueGramX, uniqueGramY))
	fmt.Println("差集合", difference(uniqueGramX, uniqueGramY))
	fmt.Println("Xにseが含まれるか", isInclude(biGramX, "se"))
	fmt.Println("Yにseが含まれるか", isInclude(biGramY, "se"))

/* =>
   X bi-gram [pa ar ra ap pa ar ra ap pa ar ra ad di is se]
   Y bi-gram [pa ar ra ag gr ra ap ph]
   和集合 [pa ar ra ap ad di is se ag gr ph]
   積集合 [pa ar ra ap]
   差集合 [ad di is se]
   Xにseが含まれるか true
   Yにseが含まれるか false
*/
}
