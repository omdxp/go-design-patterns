package main

import "strings"

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type User2 struct {
	names []uint8
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i, v := range allNames {
			if v == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	parts := strings.Split(fullName, " ")
	result := User2{}
	for _, v := range parts {
		result.names = append(result.names, getOrAdd(v))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, v := range u.names {
		parts = append(parts, allNames[v])
	}
	return strings.Join(parts, " ")
}

func main() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane Smith")
	totalMem := 0
	totalMem += len(john.FullName)
	totalMem += len(jane.FullName)
	totalMem += len(alsoJane.FullName)
	println(totalMem)
	println(john.FullName)
	println(jane.FullName)
	println(alsoJane.FullName)

	john2 := NewUser2("John Doe")
	jane2 := NewUser2("Jane Doe")
	alsoJane2 := NewUser2("Jane Smith")
	totalMem = 0
	for _, v := range allNames {
		totalMem += len([]byte(v))
	}
	totalMem += len(john2.names)
	totalMem += len(jane2.names)
	totalMem += len(alsoJane2.names)
	println(totalMem) // saved 4 bytes
	println(john2.FullName())
	println(jane2.FullName())
	println(alsoJane2.FullName())
}
