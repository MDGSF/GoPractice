package main

import "person"
import "fmt"

func main() {
	p := new(person.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
