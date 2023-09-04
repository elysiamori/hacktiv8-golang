package main

import "fmt"

// func main() {
// 	fmt.Println("THIS LINE ONWARDS IS FROM MAIN FUNCTION")

// 	// visitBack()
// 	// sctFunction()
// 	// sctClosure()
// 	// sctStruct()
// 	// sctExport()
// 	jsonExample()
// }

func visitBack() {
	map1 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	array1 := []string{"1", "2", "3"}
	array2 := append([]string{}, array1...)
	array1[0] = "bukan 1"
	fmt.Println(array2)

	for key, val := range map1 {
		fmt.Printf("key:%s ", key)
		fmt.Printf("val:%d\n", val)
	}
}
