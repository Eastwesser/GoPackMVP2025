package main

import "fmt"

func commentInQuotationMarks(quotationMark string) string {

	if quotationMark == "`" {
		fmt.Println("[``` используй такие кавычки (Ё или ~) ```]")
		return "```"
	}

	return quotationMark
}

func main() {
	quote := "`"

	mark := commentInQuotationMarks(quote)
	fmt.Println(mark)
}
