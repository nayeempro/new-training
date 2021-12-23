package main

import "fmt"

func main() {

	/*
		fmt.Print("Enter your age: ")
		var age int
		fmt.Scanf("%d", &age)

		//if boolean expression {}

			if age < 3 {

				fmt.Println("infant")
			} else if age > 2 && age < 12 {

				fmt.Println("children")

			} else if age > 2 && age < 12 {

				fmt.Println("children")

			} else if age > 12 && age < 19 {

				fmt.Println("teen age")

			} else if age > 19 && age < 25 {

				fmt.Println("adult")
			}
	*/

	//For loop
	/*
		for i := 1; i <= 9; i++ {

			fmt.Println(i)

		}
	*/

	/*
		students := []string{"Asgor", "al", "bl", "cl"}
		for i, stu := range students {
			fmt.Println(i, stu)
		}
	*/
	i := 1
	for i <= 15 {
		fmt.Println(i, "ok")
		i++
	}

}
