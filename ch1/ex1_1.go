//Exercise 1.1

//Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
  "fmt"
  "os"
)

//This is the inefficient, looping version of this program. See ex1_3 for the more efficient one using strings.Join.

func main(){
  var arguments string

  //Adds os.Args[0] to the beginning of the arguments string.
  arguments += os.Args[0]

  //Lopos over the other arguments and adds them to the string.
  for _, arg := range os.Args[1:] {
    arguments += " " + arg
  }

  //Prints the string.
  fmt.Println(arguments)
}
