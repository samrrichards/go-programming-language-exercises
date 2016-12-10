//Exercise 1.3

//Experiment to measure the difference in running time between our potentially inefficient
//versions and the one that uses strings.Join.

package main

import (
  "fmt"
  "os"
  "strings"
)

//Here is the more efficient version of the program in ex1_1:

func main() {
  fmt.Println(strings.Join(os.Args[0:], " "))
}

//
