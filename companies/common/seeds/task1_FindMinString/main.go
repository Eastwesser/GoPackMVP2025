package main

import "fmt"

//func FindMax[T comparable](slice []T, compare func(a, b T) bool) (T, error) {
//	if len(slice) == 0 {
//		var zero T
//		return zero, fmt.Errorf("slice is empty")
//	}
//
//	max := slice[0]
//	for i := 1; i < len(slice); i++ {
//		if compare(slice[i], max) {
//			max = slice[i]
//		}
//	}
//	return max, nil
//}
//
//// Использование:
//maxInt, _ := FindMax([]int{3, 1, 4, 2}, func(a, b int) bool { return a > b })
//longestStr, _ := FindMax([]string{"a", "bb", "ccc"}, func(a, b string) bool { return len(a) > len(b) })

// Max for ints
func FindMaxInt(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("Cannot find the maximum number of ints/slice is empty")
	}
	maximal := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > maximal {
			maximal = slice[i]
		}
	}

	return maximal, nil
}

func FindMinInt(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("Cannot find the minimum number of ints/slice is empty")
	}
	minimal := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < minimal {
			minimal = slice[i]
		}
	}
	return minimal, nil
}

func FindMaxString(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("Cannot find the maximum number of strings/slice is empty")
	}
	maximal := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > maximal {
			maximal = slice[i]
		}
	}
	return maximal, nil
}

func FindMinString(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("Cannot find the minimum number of strings/slice is empty")
	}
	minimal := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < minimal {
			minimal = slice[i]
		}
	}
	return minimal, nil
}

func medic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from:", err)
		}
	}()
	fmt.Println("Hello World with panic")
}

/*
	[T any] - параметр типа (дженерик)
	func(a, b T) bool - функция сравнения
	var zero T - zero value для типа T
*/

func FindMin[T any](slice []T, compare func(a, b T) bool) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, fmt.Errorf("Cannot find the minimum number of ints/slice is empty")
	}

	minimal := slice[0]
	for i := 1; i < len(slice); i++ {
		if compare(slice[i], minimal) {
			minimal = slice[i]
		}
	}
	return minimal, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stringSlice := []string{"A", "B", "C", "D", "E", "F"}

	maxInt, err := FindMaxInt(intSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxInt)
	}

	minInt, err := FindMinInt(intSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxInt, minInt)
	}

	maxString, err := FindMaxString(stringSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxString, stringSlice)
	}

	minString, err := FindMinString(stringSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxString, minString)
	}

	// для дженерика
	minString2, err := FindMin(stringSlice, func(a, b string) bool { return a < b })
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(minString2, stringSlice)
	}

}
