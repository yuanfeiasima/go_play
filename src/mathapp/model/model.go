package model

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}



