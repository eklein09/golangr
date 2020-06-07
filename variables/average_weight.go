package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter weight of first person:")
	weight1, _ := reader.ReadString('\n')
	weight1 = weight1[:len(weight1)-1]
	intWeight1, _ := strconv.Atoi(weight1)

	fmt.Println("Enter weight of second person:")
	weight2, _ := reader.ReadString('\n')
	weight2 = weight2[:len(weight2)-1]
	intWeight2, _ := strconv.Atoi(weight2)

	fmt.Println("Enter weight of third person:")
	weight3, _ := reader.ReadString('\n')
	weight3 = weight3[:len(weight3)-1]
	intWeight3, _ := strconv.Atoi(weight3)

	fmt.Println("Enter weight of fourth person:")
	weight4, _ := reader.ReadString('\n')
	weight4 = weight4[:len(weight4)-1]
	intWeight4, _ := strconv.Atoi(weight4)

	fmt.Println("Enter weight of fifth person:")
	weight5, _ := reader.ReadString('\n')
	weight5 = weight5[:len(weight5)-1]
	intWeight5, _ := strconv.Atoi(weight5)

	averageWeight := float32(intWeight1+intWeight2+intWeight3+intWeight4+intWeight5) / 5.0

	fmt.Printf("Average weight is %f\n", averageWeight)

}
