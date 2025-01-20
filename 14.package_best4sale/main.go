package main

import (
	"errors"
	"fmt"
	"os"
)

func checkDiscountAvito(price, discount float64) (float64, error) {
	var result float64

	if discount > 90.0 {
		return 0, errors.New("the sale cannot exceed 90%")
	}
	result = price * (1 - discount/100)

	fmt.Println(result)
	return result, nil
}

func main() {
	price := 1999.0
	discount := 30.0

	discountedPriceAvito, err := checkDiscountAvito(price, discount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your price, sale included:", discountedPriceAvito)
	}

	os.Exit(0)
}
