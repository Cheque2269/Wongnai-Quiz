package main

import (
	"encoding/base64"
	"fmt"
)

func reverseString(word string) string {
	str := []rune(word)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = reverseString(string(sd))

	fmt.Println(whatIsIt)
}
