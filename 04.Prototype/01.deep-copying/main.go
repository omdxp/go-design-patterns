package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}}
	// jane := john
	// jane.Name = "Jane"
	// jane.Address.StreetAddress = "321 Baker Street"

	// fmt.Println(john, john.Address)
	// fmt.Println(jane, jane.Address)

	// Deep Copying
	jane2 := john
	jane2.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane2.Name = "Jane"
	jane2.Address.StreetAddress = "321 Baker Street"

	fmt.Println(john, john.Address)
	fmt.Println(jane2, jane2.Address)
}

// Deep Copying is a process of creating a new object with the same value as an existing object.
