package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Please enter a number between 1 and 10: ")
    number, _ := reader.ReadString('\n')
    number = number[:len(number)-1]
    int_number, _ := strconv.Atoi(number)
    if (int_number <= 10 && int_number >=1) {
        fmt.Println("The number was between 1 and 10")
    } else {
        fmt.Println("Number was not between 1 and 10")
    }
    }
