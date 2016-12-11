//Exercise 1.1

//Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
  "fmt"
  "os"
  "strings"
)

//Prints a string containing everything passed in to os.Args, separated by spaces.

func main() {
  fmt.Println(strings.Join(os.Args[0:], " "))
}
