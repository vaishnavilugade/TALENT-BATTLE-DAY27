package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number:\n")
	fmt.Scan(&num)
	temp := num
	for i := 1; i <= num; i++ {
		temp++

	}
	fmt.Print("Double of ", num, " is ", temp)
}
