package model

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type People struct {
	People []Person `json:"data"`
}
