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
}
