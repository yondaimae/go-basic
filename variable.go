package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*Format1 */
	fmt.Println("<========Declare Variable========>")
	var strName, strAddress, strCompany string
	var number int
	strName, strAddress, strCompany = "My name", "Address", "N-digital"
	number = 100
	fmt.Printf("Name:%s \n Address:%s\n Company:%s\n Number:%d\n", strName, strAddress, strCompany, number)

	/*Format2 */
	var strNameF2, strAddressF2, strCompanyF2, numberF2 = "My name F2", "Address F2", "N-digital F2", 200
	fmt.Printf("NameF2:%s \n AddressF2:%s\n CompanyF2:%s\n NumberF2:%d\n", strNameF2, strAddressF2, strCompanyF2, numberF2)

	/*Format3 */
	var strNameF3, strAddressF3, strCompanyF3 string = "My name F3", "Address F3", "N-digital F3"
	var numberF3 int = 300
	fmt.Printf("NameF3:%s \n AddressF3:%s\n CompanyF3:%s\n NumberF3:%d\n", strNameF3, strAddressF3, strCompanyF3, numberF3)

	/*Format4
	ประกาศตัวแปรใหม่ใช้ได้เฉพาะภายในฟังก์ชั่น
	*/
	strNameF4, strAddressF4, strCompanyF4, numberF4 := "My name F4", "Address F4", "N-digital F4", 400
	fmt.Printf("NameF4:%s \n AddressF4:%s\n CompanyF4:%s\n NumberF4:%d\n", strNameF4, strAddressF4, strCompanyF4, numberF4)

	/*Array Format1*/
	fmt.Println("<========Array========>")
	var arrayF1 [2]string
	arrayF1[0] = "Hello"
	arrayF1[1] = "Golang"
	fmt.Printf("Array F1:%s\n", arrayF1)

	/*Array Format2*/
	arrayF2 := [3]int{1, 2, 3}
	fmt.Printf("Array F2:%d\n", arrayF2)

	/*Array Format3
	Multiple Dimension
	*/
	arrayF3 := [3][3]int{{1}, {1, 2}, {1, 2, 3}}
	fmt.Printf("Array F3:%d\n", arrayF3)

	/*Slice
	Array Dynamic Size
	*/
	fmt.Println("<========Slice========>")
	sliceF1 := make([]int, 0)
	sliceF1 = append(sliceF1, 1)
	fmt.Printf("Slice F1:%d\n", sliceF1)

	/*Slice Format5
	Multiple Dimension and dynamic size
	Dynamic slice short x := []int{}
	*/
	sliceF2 := make([][]int, 0)
	sliceF2 = append(sliceF2, []int{0, 1, 2})
	sliceF2 = append(sliceF2, []int{3, 4, 5})
	sliceF2 = append(sliceF2, []int{6, 7, 8})
	fmt.Printf("Slice F2:%d\n", sliceF2)

	/*Example Slice Pointer*/
	fmt.Println("<========Slice Pointer========>")
	mirrorSliceF2 := sliceF2
	mirrorSliceF2[0] = []int{9, 10, 11}
	fmt.Printf("sliceF2:%d\n", sliceF2)
	fmt.Printf("mirrorSliceF2:%d\n", mirrorSliceF2)

	/*Slice no pointer */
	fmt.Println("<========Slice No Pointer========>")
	mirrorSliceF3 := make([][]int, len(sliceF2))
	copy(mirrorSliceF3, sliceF2)
	mirrorSliceF3[1] = []int{9, 10, 11}
	fmt.Printf("sliceF2:%v\n", sliceF2)
	fmt.Printf("mirrorSliceF3:%v\n", mirrorSliceF3)

	/*Struct*/
	fmt.Println("<========Struct========>")
	type MyData struct {
		Col1 string
		Col2 int
	}
	type Message struct {
		Name   string
		Number int
		Data   interface{}
	}

	structMessage := Message{"MyName", 100, MyData{"Sample", 100}}
	fmt.Printf("Struct:%+v\n", structMessage)

	/*Json Struct*/
	fmt.Println("<========Json Struct========>")
	structJson, _ := json.Marshal(structMessage)
	fmt.Printf("Json:%s\n", structJson)
	var decodeJson interface{}
	err := json.Unmarshal(structJson, &decodeJson)
	if err != nil {
		fmt.Printf("Error decode json:%v", err)
	}
	fmt.Printf("Decode json:%+v\n", decodeJson)

	/*Json Interface*/
	fmt.Println("<========Json Interface========>")
	jsonInterface := map[string]interface{}{
		"ID":      1,
		"Message": "Help",
	}
	interfaceJson, _ := json.Marshal(jsonInterface)
	fmt.Printf("Json interface:%s\n", interfaceJson)

	/*Map*/
	fmt.Println("<========Map========>")
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["xxx"] = 3
	fmt.Println("Numbers:", numbers)

	/*Map Interface*/
	fmt.Println("<========Map========>")
	mapArr := []map[string]interface{}{
		{
			"ID":   1,
			"Name": "Sample",
		},
		{
			"ID":   2,
			"Name": "Sample2",
		},
		{
			"ID": 3,
			"Name": map[string]interface{}{
				"Extra": 1,
			},
		},
		{
			"ID": 4,
			"Name": []map[string]interface{}{
				{
					"Extra": 1,
				},
				{
					"Extra": 2,
				},
			},
		},
	}

	fmt.Println("MapArr:", mapArr)

	/*Print format
	reference:https://siamchamnankit.co.th/%E0%B8%88%E0%B8%94%E0%B9%84%E0%B8%A7%E0%B9%89%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%A5%E0%B8%B7%E0%B8%A1-%E0%B8%A7%E0%B8%B1%E0%B8%99%E0%B9%81%E0%B8%A3%E0%B8%81%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%84%E0%B8%A5%E0%B8%B2%E0%B8%AA-basic-golang-part-ii-type-array-slice-map-and-struct-ac0c89e3b725
	reference:https://pkg.go.dev/fmt
	reference:https://medium.com/rungo/string-formatting-in-go-dd752ff7f60
	*/
}
