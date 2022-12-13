package main

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployeeFactory(role int) func(name string) *Employee {
	switch role {
	case Developer:
		return func(name string) *Employee {
			return &Employee{name, "Developer", 60000}
		}
	case Manager:
		return func(name string) *Employee {
			return &Employee{name, "Manager", 80000}
		}
	default:
		panic("unsupported role")
	}
}

func main() {
	developerFactory := NewEmployeeFactory(Developer)
	managerFactory := NewEmployeeFactory(Manager)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	println(developer.Name, developer.Position, developer.AnnualIncome)
	println(manager.Name, manager.Position, manager.AnnualIncome)
}

// prototype factory is a factory that returns a function that returns a pointer to a struct
// it is used to create a new instance of a struct
// it is used to hide the implementation details of the struct
// it is used to create a new instance of a struct with some default values
