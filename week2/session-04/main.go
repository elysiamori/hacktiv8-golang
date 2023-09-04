package main

import (
	"fmt"
	"reflect"
	"session-04/model"
	"sync"
)

type TestIota uint16

const (
	Const1 TestIota = iota + 1
	Const2
	Const3
)

func refresh() {
	p1 := model.Person{
		Name: "Alam",
	}

	fmt.Println("Test Iota", Const1, Const2, Const3)

	fmt.Println(p1)
}

func main() {
	// refresh()
	// sctReflect()
	// sctPointer()
	// sctInterface()
	// sctEmptyInterface()
	// sctConcurrency()
	additional()
}

type Student struct {
	Name       string `testTag:"test1" testTag2:"test2"`
	Score      uint16
	pesanKesan string
}

func (s Student) String() string {
	return fmt.Sprintf(`
	Nama: %s
	Nilai: %d
	Pesan Kesan: %s
	`, s.Name, s.Score, s.pesanKesan)
}

type ErrorStudent struct {
	Name string
}

func (e ErrorStudent) Error() string {
	return fmt.Sprintf(`%s melakukan pelanggaran`, e.Name)
}

func (s *Student) SetPesanKesan(input string) error {
	kataKotor := map[string]bool{
		"kata-kotor1": true,
		"kata-kotor2": true,
		"kata-kotor3": true,
		"kata-kotor4": true,
		"kata-kotor5": true,
		"kata-kotor6": true,
	}
	if _, ok := kataKotor[input]; ok {
		e := ErrorStudent{
			Name: s.Name,
		}
		return e
	}
	s.pesanKesan = input
	return nil
}

func (s *Student) getFieldsInfo() {
	fmt.Println("BELOW IS RESULT OF GETFIELDSINFO")
	valS := reflect.ValueOf(s)

	if valS.Kind() == reflect.Ptr {
		valS = valS.Elem()
	}

	typeS := valS.Type()

	for i := 0; i < valS.NumField(); i++ {
		fmt.Println("name :", typeS.Field(i).Name)
		fmt.Println("data type :", typeS.Field(i).Type)
		fmt.Println("value in interface form :", valS.Field(i).Interface())
		fmt.Println("tag :", typeS.Field(i).Tag)
	}
}

func (s *Student) SetName(prefix, newName string) (isSuccess bool) {
	s.Name = prefix + "-" + newName
	return true
}

func testWithPointer(input *int, newValue int) {
	*input = newValue
}

func testWithoutPointer(input int, newValue int) {
	input = newValue
}

func sctReflect() {
	var1 := uint16(30)

	valVar1 := reflect.ValueOf(var1)
	type1Var1 := valVar1.Type()
	type2Var1 := reflect.TypeOf(var1)

	fmt.Println(type1Var1, type2Var1)

	if valVar1.Kind() == reflect.Uint16 {
		fmt.Println(
			valVar1.Uint(),
		)
	}

	varFunc := func(a, b string) bool {
		return false
	}

	valVarFunc := reflect.ValueOf(varFunc)

	fmt.Println(valVarFunc.Type())

	if valVarFunc.Kind() == reflect.Func {
		fmt.Println("this is a function")
	}

	fmt.Println("value for var1:", valVar1.Interface())
	fmt.Println("value for valVarFunc:", valVarFunc.Interface())

	valVar1Ret := valVar1.Interface().(uint16)
	fmt.Println(valVar1Ret)

	valVarFuncRet, ok := valVarFunc.Interface().(func(a, b string))
	if ok {
		fmt.Println("\"value\" of valVarFunc", valVarFuncRet)
	}

	student1 := &Student{
		Name:  "Dominique",
		Score: 90,
	}

	student1.getFieldsInfo()

	var method = reflect.ValueOf(student1).MethodByName("SetName")
	res := method.Call([]reflect.Value{
		reflect.ValueOf("NM"),
		reflect.ValueOf("Dominique Jeffrey"),
	})

	fmt.Println("========================")
	fmt.Println(res[0].Interface())
	fmt.Println("========================")
	student1.getFieldsInfo()
}

func sctPointer() {
	num1 := float64(14)
	num2 := &num1

	fmt.Println(num1, *num2)
	fmt.Println(&num1, num2)

	num3 := 70
	testWithoutPointer(num3, 71)
	fmt.Println(num3)
	testWithPointer(&num3, 73)
	fmt.Println(num3)

	var name1 *Student = new(Student)
	name1.Name = "alam"
	fmt.Println(name1.Name)
}

type Animal interface {
	Move() string
	Incubation() uint // in weeks
	RelativeLifeSpan() float32
	// DefenseMechanism()
}

type Dog struct {
	Age uint
}

func (d Dog) Move() string {
	return "walking / running"
}

func (d Dog) Incubation() uint {
	return 12
}
func (d Dog) RelativeLifeSpan() float32 {
	return float32(d.Age) / 20
}

func (d Dog) DefenseMechanism() {
	d.Barking()
}

func (d Dog) Barking() {
	fmt.Println("the dog is barking")
}

type Bird struct {
	Age uint
}

func (b Bird) Move() string {
	return "flying"
}

func (b Bird) Incubation() uint {
	return 8
}

func (b Bird) RelativeLifeSpan() float32 {
	return float32(b.Age) / 10
}

type Fish struct {
	Age uint
}

func (f Fish) Move() string {
	return "swimming"
}

func (f Fish) RelativeLifeSpan() float32 {
	return float32(f.Age) / 15
}

func (f Fish) Incubation() uint {
	return 8
}

func compareAnimalRelativeAge(animal1, animal2 Animal) {
	if animal1.RelativeLifeSpan() >= animal2.RelativeLifeSpan() {
		fmt.Println("animal 1 is as old or older than animal 2")
		return
	}

	fmt.Println("animal 2 is older than animal 1")

}

func sctInterface() {
	var dog Animal = Dog{
		Age: 5,
	}
	var bird Animal = Bird{
		Age: 8,
	}
	var fish Animal = Fish{
		Age: 3,
	}

	realDog := dog.(Dog)
	fmt.Println(dog, bird, fish, realDog)
	realDog.Barking()

}

func sctEmptyInterface() {
	var anything interface{}

	anything = 1
	anything = "1"
	anything = Bird{}
	anything = 1

	anythingFormInt := anything.(int)
	anythingFormInt = anythingFormInt + 1

	s := Student{
		Name:  "Dominique",
		Score: 8,
	}

	err := s.SetPesanKesan("semuanya baik")
	if err != nil {
		fmt.Println(err)
	}

	err = s.SetPesanKesan("kata-kotor1")
	if err != nil {
		fmt.Println(err)
	}
}

func RunWithGoroutine() {
	fmt.Println("this is using goroutine")
}

func sctConcurrency() {
	// A:0-3; 5-7
	// B:0-2; 4-5
	// (Serial) Thread: {A:0-3; 5-7} {B:7-9; 11-12}
	// (Concurrency) Thread: {A:0-3}{B:3-5} {A:5-7} {B:7-8}

	// go RunWithGoroutine()
	// fmt.Println("this is main goroutine")
	// fmt.Println(runtime.NumGoroutine())
	// time.Sleep(10)

	students := []Student{
		{
			Name:  "Dominique",
			Score: 8,
		},
		{
			Name:  "Jeffrey",
			Score: 9,
		},
		{
			Name:  "Alamaro",
			Score: 10,
		},
	}
	var wg sync.WaitGroup

	for _, student := range students {
		s := student
		wg.Add(1)
		go func() {
			// time.Sleep(time.Second * 10)
			fmt.Println(s.Name, &wg)
			wg.Done()
		}()
	}

	wg.Wait()

}

func additional() {
	var num int = 40
	valNum := reflect.ValueOf(num)
	fmt.Printf("%s, %d\n", valNum.Kind(), valNum.Kind())

	s1 := Student{
		Name:       "dom",
		Score:      9,
		pesanKesan: "something",
	}
	fmt.Println(s1)
}

func Foo(name interface{}) {
	st, ok := name.(fmt.Stringer)
	if !ok {
		// pakai default golang
	}

	st.String()
}

type HousingRepo struct {
}

func (h HousingRepo) ProcessEndpoint1() {
	// yang berhubungan dengan database & datasources

}

type HousingService struct {
	housingRepo HousingRepo // biasanya bukan struct tapi interface
}

func (h HousingService) ProcessEndpoint1() {
	// handle business logic

	h.housingRepo.ProcessEndpoint1()
}

type HousingController struct {
	housingServ HousingService // biasanya bukan struct tapi interface
}

func (h HousingController) Endpoint1() {
	// bla-bla yang berhubungan dengan request user

	h.housingServ.ProcessEndpoint1()
}
