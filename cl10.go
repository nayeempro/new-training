package main

import "fmt"

func main() {
	/*var name [3]string
	name[0] = "Nicola"
	name[1] = "Murad"
	name[2] = "Nahid raince"
	fmt.Println(name)
	fmt.Println(name[1])
	*/
	//short way
	//... means joto khusi nite paro
	//name := [...]string{"Takla Murad", "Toslima Nasrin", "Murgi kobir", "hasao mahmud"}
	//fmt.Println(name)
	// fmt.Println(name[2])
	//x := name[0:4]
	//fmt.Println(x)
	/*var fruits []string
	fruits = append(fruits, "apple", "Banana", "Mango")
	fmt.Println(fruits)
	fmt.Printf("%T", fruits)  */

	/*
		var x map[string]int
		x["Nayeem"] = 6
		fmt.Println(x)

	*/

	x := make(map[string]int)
	x["Nayeem"] = 3
	x["Rashid"] = 5
	x["Rakib"] = 5

	fmt.Println(x["Nayeem"])
	fmt.Println(x)

}
