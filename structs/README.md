// basic of structs

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// lazy syntax, could lead to errors if the struct changes
	// depends on the field order on the struct definition
	// alex := person{"Alex", "Anderson"}

	// clearer and safe way
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	// another way, by declaring the var with zero value
	var alex_ person
	alex_.firstName = "Alex"
	alex_.lastName = "Anderson"

	fmt.Println(alex_)
	// the %+v shows the content of the structs with the its fields
	fmt.Printf("%+v", alex_)
}


// pointers

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// don't need to use a identifier for structs inside a struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halper",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// & operator returns the memory address of a variable
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

// * operator in front of a type means it expect a pointer to the type
func (pointerToPerson *person) updateName(newFirstName string) {
	// * operator in front of a var returns the value of the memory address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
