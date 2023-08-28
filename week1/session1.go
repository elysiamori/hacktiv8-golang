package main

import "fmt"

func main() {
    fmt.Println("hello world")
    var namaVariable string
    fmt.Println("ini nilai dari variable kita -" + namaVariable + "-")

    fmt.Printf("hello %s", "alam\n")

    // var decimal float32=1.999
    decimal := 1.99999999999

    decimal = 11111
    fmt.Printf("%.10f", decimal)

    // var (
    //     var1, var2, var3 string
    //     // var4             float32
    // )

    // var1, var2, var3 = "variable1", "variable2", "variable3"

    _, _, _ = "string", 1, 1.123123

    slice := []string{"array1", "array2", "array3"}
    for _, el := range slice {
        fmt.Println(el)
    }

    test := new(string) // pointer to a string variable with the value of "" (default value for string)
    fmt.Println("-" + *test + "-")

    testRune := "ini adalah string"
    for _, el := range testRune {
        fmt.Printf("%s", string(el))
    }
    fmt.Println()

    var varBool bool = true
    fmt.Printf("%t\n", varBool)
    cobaMultiline := `
        alam \n alam \n alam \n
        " ' 
    `
    fmt.Println(cobaMultiline)

    var testNil *string = new(string)
    fmt.Println(testNil, "-"+*testNil+"-")

    var testFunction func(string)
    testFunction = func(s string) {
        fmt.Println(s)
    }

    testFunction("hello, world")

    const constant1 string = "THIS IS OUR FIRST CONSTANT"

    // 8/3 -> 2.67
    // 8/3 -> 2 sisa 3
    // testNumber := 3
    var testDiv float64 = float64(7) / float64(4)
    fmt.Println(testDiv)

    varTrue := true
    varFalse := false

    fmt.Printf("%t %tn", varTrue, varFalse)
}
``
