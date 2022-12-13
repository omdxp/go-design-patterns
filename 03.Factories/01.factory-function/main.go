package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 0 {
		age = 0
	}
	return &Person{Name: name, Age: age, EyeCount: 2}
}

func main() {
	p := NewPerson("John", 20)
	println(p.Name, p.Age, p.EyeCount)
}

// factory function is a function that returns a pointer to a struct
// it is used to create a new instance of a struct
// it is used to hide the implementation details of the struct
// it is used to create a new instance of a struct with some default values
