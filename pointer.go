package main

import (
	"fmt"
	"strings"
)

func changeVal(a int) {
	fmt.Println("<=====Chang Value=====>")
	fmt.Printf("Address a:%v\n", &a)
	fmt.Printf("Value a:%v\n", a)

}
func changeValPointer(a *int) {
	fmt.Println("<=====Chang Pointer Value=====>")
	fmt.Printf("Address a:%v\n", a)
	fmt.Printf("Value a:%v\n", *a)
}
type Profile struct {
	Name string
	Age  int
}

func upperProfileName(p *Profile) {
	p.Name = strings.ToUpper(p.Name)
}
func main() {
	/*Pointer
	reference:https://dev.to/iporsut/go-pointer-pointer-go-3212
	*/
	var a int
	fmt.Println("<=====Main Value=====>")
	fmt.Printf("Address value a:%v\n", &a)
	fmt.Printf("Value a:%v\n", a)
	changeVal(5)
	fmt.Println("<=====Main Value=====>")
	fmt.Printf("Address value a:%v\n", &a)
	fmt.Printf("Value a:%v\n", a)

	pa := &a
	*pa = 5
	changeValPointer(pa)
	fmt.Println("<=====Main Value=====>")
	fmt.Printf("Address value a:%v\n", &a)
	fmt.Printf("Value a:%v\n", a)

	fmt.Println("<=====Pointer of struct=====>")
	p := Profile{
		Name: "Por",
		Age:  35,
	}
	upperProfileName(&p)
	fmt.Println(p.Name)

}
