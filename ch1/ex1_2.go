//Exercise 1.1

//Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
  "fmt"
  "os"
)

//Loops over the arguments passed in and prints out their index number and value.

func main(){
  for index, arg := range os.Args[1:] {
    fmt.Println(index, arg)
  }
}
