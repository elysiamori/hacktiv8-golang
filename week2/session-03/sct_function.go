package main

import "fmt"

func sctFunction() {
	greetRes := greet("dominique", "bekasi", 26)
	fmt.Println(greetRes)

	area1, oL1 := calculateSegiPanjang(10.5, 8.6)
	fmt.Printf("area is %.2f and outer line is %.2f\n", area1, oL1)

	fmt.Println(calculateTotal(1, 2, 3, 4, 5, 6))

	students := []string{"dominqiue", "jeffrey", "alamaro", "maximilianus"}
	fmt.Println(processStudents(students...))

	printTodayActivities("dominique", "worked", "played", "ate")

	input := "dominique"
	testParseAndModify(input)
	fmt.Println("after testParseAndModify", input)

	// input2 := "dominique2"
	testParseAndModify2(&input)
	fmt.Println("after testParseAndModify2", input)

	inputMap := map[string]string{}
	testParseAndModify3(inputMap)
	fmt.Println(inputMap)

}

func greet(name, address string, age uint8) string {
	return fmt.Sprintf("hello %s with address of %s with age of %d\n", name, address, age)
}

func calculateSegiPanjang(length, width float64) (area float64, outerLine float64) {
	area = length * width
	outerLine = 2 * (length + width)
	return area, outerLine
}

func calculateTotal(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func processStudents(students ...string) (res map[string]string) {
	res = map[string]string{}

	for i, student := range students {
		res[fmt.Sprintf("student%d", i+1)] = student
	}

	return
}

func printTodayActivities(name string, activites ...string) {
	fmt.Println("Today " + name + "did this list below:")
	for _, act := range activites {
		fmt.Println(act)
	}
}

func testParseAndModify(input string) {
	input = fmt.Sprintf(`%s next version`, input)
}

func testParseAndModify2(input *string) {
	*input = fmt.Sprintf(`%s next version`, *input)
}

func testParseAndModify3(input map[string]string) {
	input["testParseAndModify3"] = "testParseAndModify3 value"
}
