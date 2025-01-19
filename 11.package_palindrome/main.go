package main

import (
	"fmt"
	"strings"
)

func ifIsPalindrome(str string) bool {
	var str1, str2 string

	str = strings.ToLower(str)

	//for i := 0; i < len(str); i++ {
	//	if str[i] != str[len(str)-i-1] {
	//		return false
	//	}
	//}

	for i := 0; i < len(str); i++ {
		str1 = string(str1[i])
	}

	for i := len(str); i >= 0; i-- {
		str2 = string(str2[i])
	}

	if str1 == str2 {
		return true
	}
	return false
}

func main() {
	var newWord string
	fmt.Scanln(&newWord)
	fmt.Println(ifIsPalindrome(newWord))
}

//PS C:\Users\altte\OneDrive\Desktop\GoPackMVP2025\11.package_palindrome> go run main.go
//radar
//panic: runtime error: index out of range [0] with length 0
//
//goroutine 1 [running]:
//main.ifIsPalindrome({0xc00000a048?, 0xc0000900b0?})
//C:/Users/altte/OneDrive/Desktop/GoPackMVP2025/11.package_palindrome/main.go:20 +0x2f
//main.main()
//C:/Users/altte/OneDrive/Desktop/GoPackMVP2025/11.package_palindrome/main.go:36 +0x6c
//exit status 2
//PS C:\Users\altte\OneDrive\Desktop\GoPackMVP2025\11.package_palindrome>
