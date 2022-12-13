package main

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	// functional approach
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	println(developer.Name, developer.Position, developer.AnnualIncome)
	println(manager.Name, manager.Position, manager.AnnualIncome)

	// structural approach
	bossFactory := EmployeeFactory{"CEO", 100000}
	bossFactory.AnnualIncome = 200000
	boss := bossFactory.Create("Sam")

	println(boss.Name, boss.Position, boss.AnnualIncome)
}

// factory generator is a function that returns a function that returns a pointer to a struct in case of a functional approach or a method that returns a pointer to a struct in case of a structural approach
// it is used to create a new instance of a struct
// it is used to hide the implementation details of the struct
// it is used to create a new instance of a struct with some default values
