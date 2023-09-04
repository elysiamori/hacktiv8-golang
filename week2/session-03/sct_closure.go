package main

import "fmt"

func sctClosure() {
	// closure as variable
	var oddNumbers = func(numbers ...int) []int {
		res := []int{}
		for _, num := range numbers {
			if num%2 == 1 {
				res = append(res, num)
			}
		}
		return res
	}

	fmt.Println("oddNumbers (the variable is a closure)", oddNumbers)

	oddArray := oddNumbers([]int{1, 2, 3, 4, 5}...)
	fmt.Println(oddArray)

	// immediately-invoked closure
	var immediatelyInvokedOddNumbers = func(numbers ...int) []int {
		res := []int{}
		for _, num := range numbers {
			if num%2 == 1 {
				res = append(res, num)
			}
		}
		return res
	}([]int{1, 2, 3, 4, 5}...)

	fmt.Println("immediatelyInvokedOddNumbers (the variable is a array of integer)", immediatelyInvokedOddNumbers)

	inputForOddAndEvenFunc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddFunc := oddOrEvenArray(true)
	evenFunc := oddOrEvenArray(false)

	fmt.Println("odd array", oddFunc(inputForOddAndEvenFunc))
	fmt.Println("even array", evenFunc(inputForOddAndEvenFunc))

	callback := func(in int) {
		fmt.Println(in * in * in)
	}
	printSomethingWithNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, callback)
}

// closure as output
func oddOrEvenArray(isOdd bool) func([]int) []int {
	if isOdd {
		return func(numbers []int) []int {
			res := []int{}
			for _, num := range numbers {
				if num%2 == 1 {
					res = append(res, num)
				}
			}
			return res
		}
	}

	return func(numbers []int) []int {
		res := []int{}
		for _, num := range numbers {
			if num%2 == 0 {
				res = append(res, num)
			}
		}
		return res
	}

}

// closure as input (callback)
func printSomethingWithNumbers(numbers []int, cb func(int)) {
	for _, num := range numbers {
		fmt.Printf("executing callback of %d \t", num)
		cb(num)
	}
}
