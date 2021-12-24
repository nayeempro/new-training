package main

import "fmt"

//method 1
/*
func add(x int, y int) int {

	r := x + y

	return r
}
*/
//method 2  ekoi data type er hole
/*
func add(x, y int) int {

	r := x + y

	return r
}
*/
// method 3 return take age define korechi
/*
func add(x, y int) (r int) {

	r = x + y

	return r
}
*/
//method 4
/*
func add(x, y int) (r int) {

	r = x + y

	return
}
*/
//method 5
/*
func add(x, y int) (r int, p int) {

	r = x + y
	p = x * y

	return
}
*/
//pointer
/*
func update(x *int, y *string) {

	*x = *x + 8
	*y = *y + "Nayeem"

	return
}
*/

func main() {

	//ad, mul := add(10, 30)

	//fmt.Println(ad, mul)

	//pointer
	/*
		add := 10
		name := "Md "
		update(&add, &name)
		fmt.Println(add, name)
	*/
	//annynomus function
	a := func(x, y int) (r int) {

		r = x + y
		return

	}(20, 20)

	//fmt.Println(a(20, 30))
	fmt.Println(a)
}
