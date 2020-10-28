package main

import (
	"fmt"
	// "./mymath"
)

//Student strut
type Student struct {
	name string
}

// SetName set student name
func SetName(std *Student) {
	std.name = "Tadeyo"
}

// GetStudent  return student
func GetStudent() (std Student) {
	SetName(&std)
	return
}

func main() {
	x := m.sqrt(4)
	// fmt.Print(mymath.sqrt(4))
	// fmt.Print("Hello there")
	// fmt.Print("Hello Hello")
	fmt.Printf("%v", GetStudent())

}
