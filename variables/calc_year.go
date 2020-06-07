package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your birth year:")
	birthYear, _ := reader.ReadString('\n')

	birthYear = birthYear[:len(birthYear)-1]
	int_birthYear, _ := strconv.Atoi(birthYear)

	fmt.Println("Enter your age:")
	age, _ := reader.ReadString('\n')

	age = age[:len(age)-1]
	int_age, _ := strconv.Atoi(age)
	fmt.Printf("Entered birth year was %d\n", int_birthYear)
	fmt.Printf("Entered age was %d\n", int_age)

	currentYear := int_birthYear + int_age
	fmt.Printf("The current year is %d\n", currentYear)

}
