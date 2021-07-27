package main

import "fmt"

func main() {

	/*Interface*/
	var dynamicTypeVal interface{} = 1
	fmt.Printf("Interface integer:%T\n", dynamicTypeVal)
	dynamicTypeVal = false
	fmt.Printf("Interface boolean:%T\n", dynamicTypeVal)
	dynamicTypeVal = nil
	fmt.Printf("Interface nil:%T\n", dynamicTypeVal)
	dynamicTypeVal = "I'm String"
	fmt.Printf("Interface string:%T\n", dynamicTypeVal)
	//var strDynamicType string = dynamicTypeVal
	//fmt.Printf("strDynamicType string:%T\n", strDynamicType)
}
