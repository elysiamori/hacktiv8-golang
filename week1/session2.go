package main

import "fmt"

func main() {
	// commend-uncommend part according to what you want to see
	// conditions()
	// loopings()
	// array()
	// slice()
	// maps()
	// aliase()
}

func conditions() {
	/*
		if (boolean) -> do something
	*/

	// basic
	x := 8
	if x > 7 {
		fmt.Printf("%d lebih besar dari 7 \n", x)
	}

	// basic but condition is variable
	y := x > 7
	if y {

	}

	currYear := 2023
	bornYear := 1997
	if age := currYear - bornYear; age > 17 {
		fmt.Println("Eligible to hold KTP")
		fmt.Println(age)
	} else if age < 14 {
		fmt.Println("normally is below senior high school")
	} else {
		fmt.Println("currently studying at senior high school")
	}

	/*
		SWITCH
	*/

	var score = 6

	switch score {
	case 8:
		fmt.Println("this variable is 8")
	case 10:
		fmt.Println("this variable is 10")
	default:
		fmt.Println("this variable is neither 8 nor 10")
	}

	switch {
	case score < 8 && score > 3:
		fmt.Println("this variable is between 3 and 8")
		fallthrough
	case score == 10:
		fmt.Println("this variable is 10")
	}

	// nested if
	mathScore := 9
	scienceScore := 7

	if mathScore >= 7 {
		if scienceScore >= 7 {
			fmt.Println("ga remedial")
		} else {
			fmt.Println("remedial science")
		}
	} else {
		fmt.Println("remedial math")
	}

	// case
	const (
		COUNTRY_USA = "USA"
		COUNTRY_UK  = "UK"
		COUNTRY_INA = "INDONESIA"
		COUNTRY_IN  = "INDIA"
		COUNTRY_SG  = "SINGAPORE"
	)

	fmt.Println("with switch syntax")
	country := "INDONESIA"
	switch {
	case country == COUNTRY_INA:
		fmt.Println("selamat pagi")
	case country == COUNTRY_USA:
		fmt.Println("Good morning")
	case country == COUNTRY_UK:
		fmt.Println("Good morning tapi UK")
	case country == COUNTRY_IN:
		fmt.Println("selamat pagi dalam bahasa india")
	case country == COUNTRY_SG:
		fmt.Println("selamat pagi, ni hao, selamat pagi dalam bahasa india")
	}

	fmt.Println("with map")
	mapEnum := map[string]func(){
		COUNTRY_INA: func() { fmt.Println("selamat pagi") },
		COUNTRY_USA: func() { fmt.Println("Good morning") },
		COUNTRY_UK:  func() { fmt.Println("Good morning tapi UK") },
		COUNTRY_IN:  func() { fmt.Println("selamat pagi dalam bahasa india") },
		COUNTRY_SG:  func() { fmt.Println("selamat pagi, ni hao, selamat pagi dalam bahasa india") },
	}

	mapEnum[country]()
}

func nestedIfSolved() {
	mathScore := 9
	scienceScore := 7

	// if mathScore >= 7 {
	// 	if scienceScore >= 7 {
	// 		fmt.Println("ga remedial")
	// 	} else {
	// 		fmt.Println("remedial science")
	// 	}
	// } else {
	// 	fmt.Println("remedial math")
	// }

	if mathScore < 7 {
		fmt.Println("remedial math")
		return
	}

	if scienceScore >= 7 {
		fmt.Println("ga remedial")
	} else {
		fmt.Println("remedial science")
	}

}

func loopings() {
	/*
		for <initial state>;<looping condition>;<condition between loop>{

		}
	*/
	// "for" style 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		for <looping condition>{

		}
	*/

	// "for" style 2
	i := 0
	fmt.Println("style 2")
	for i < 10 {
		fmt.Println(i)
		i++
	}

	/*
		for{

		}
		developer needs to handle iteration and break him/herself
	*/
	// "for" style 3
	i = 0
	fmt.Println("style 3")
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	/*
		continue = keyword used to skip the rest of the current iteration
		break = break the loop
	*/
	fmt.Println("no even numbers (bilangan genap)")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("there's a number here")
			continue
		}
		fmt.Println(i)
	}

	/*
		nested looping= loop(s) in loop
	*/
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Println("i", i, "j", j)
		}
	}

	/*
		label in looping
	*/
	fmt.Println("LABEL")

loopLuar:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// if i and j greater than 5 and we find i to be even (genap) and j to be odd (ganjil)
			fmt.Println("i", i, "j", j)
			if i >= 5 && j >= 5 && i%2 == 0 && j%2 == 1 {

				break loopLuar
			}
		}
	}

}

func array() {
	var array1 [3]int
	array1 = [3]int{1, 2, 3}

	array1[1] = 10
	// fmt.Println(array1)

	// ind=index
	// el=element
	for ind, el := range array1 {
		if el >= 7 {
			array1[ind] = 7
			el = 7
		}
	}
	fmt.Println(array1)

	mulArray1 := [2][4]int{
		{1, 2, 3, 4},     //index 0
		{11, 12, 13, 14}, // index 1
	}
	fmt.Println(mulArray1)

	for k := range mulArray1 {
		for l := range mulArray1[k] {
			fmt.Println("k", k, "l", l, mulArray1[k][l])
		}
	}
}

func slice() {
	var sl1 []int = []int{0, 1, 2, 3}
	fmt.Println(sl1)

	// sl2 := make([]int, 3)

	sliceOfSl1 := sl1[:3] // kita ambil elemen dengan indeks >= bilangan dengan tapi < bilangan belakang
	fmt.Println(sliceOfSl1)

	sl3 := []int{1, 2, 3, 4, 5, 6}
	sl4 := []int{11, 12, 13}
	sl5 := append(sl3[1:3], sl4...)
	fmt.Println(sl5)

	// backing array
	fmt.Println("backing array")
	sl6 := []int{1, 2, 3}
	fmt.Println("length", len(sl6), "capacity", cap(sl6))
	sl6 = append(sl6, 1)
	fmt.Println("length", len(sl6), "capacity", cap(sl6))

	sl7 := sl6[1:]
	sl7[0] = 200
	fmt.Println("sl6", sl6, "sl7", sl7)

	sl8 := append([]int{}, sl6...)
	sl8[0] = 400
	fmt.Println("sl6", sl6, "sl8", sl8)
}

func maps() {
	/*
		map[<tipe-data-key>]<tipe-data-value>
	*/
	map1 := map[string]string{}

	map1["key-1"] = "value-1"
	// fmt.Println(map1)

	map2 := map[string]string{
		"key-1": "value-1",
		"key-2": "value-2",
	}
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, value)
	}

	map3 := map[string]string{
		"key-1": "value-1",
		"key-2": "value-2",
	}

	for key, value := range map3 {
		fmt.Println(key, value)
	}

	map4 := map[string]int{
		"key-1": 1,
		"key-2": 2,
	}

	if _, ok := map4["key-3"]; !ok {
		fmt.Println("do something")
	}

}

func aliase() {
	var a int32 = 10
	var b rune
	a = b

	fmt.Println(a, b)

	type millisecond = uint

	var mill millisecond = 10
	fmt.Printf("hour type = %T", mill)
}
