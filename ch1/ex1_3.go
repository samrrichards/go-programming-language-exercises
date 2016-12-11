//Exercise 1.3

//Experiment to measure the difference in running time between our potentially inefficient
//versions and the one that uses strings.Join.

package main

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
  "time"
)

//For some reason, the first function I call seems to take *much* longer to run than any other
//function - over 10 times longer than later function calls in some cases! The next two function
//calls also seem to take a bit longer, although by a much smaller factor.

//The last two function calls finally start producing the sorts of results the book mention on a
//consistent basis; the loop function takes about 1.5 times longer to complete than
//the join function if only a few words are entered, but this quickly jumpts to 3+ times longer as
//more words are entered. The difference in running time thus appears to be exponential.

func main() {
  loopFunc()
  joinFunc()
  loopFunc()
  joinFunc()
  loopFunc()
}

//Here is the more efficient version of the program in ex1_1:

func joinFunc() {
  defer trackTime(time.Now(), "Join")

  fmt.Println(strings.Join(os.Args[0:], " "))
}

//Here is the inefficient, looping version of the program:

func loopFunc() {
  defer trackTime(time.Now(), "Loop")

  var arguments string
  arguments += os.Args[0]

  for _, arg := range os.Args[1:] {
    arguments += " " + arg
  }

  fmt.Println(arguments)
}

//Time tracking function, which is based on ideas discussed in this blog post:
//https://blog.stathat.com/2012/10/10/time_any_function_in_go.html

func trackTime(start time.Time, name string){
  elapsed := strconv.FormatInt(time.Since(start).Nanoseconds(), 10)
  log.Printf("%s function took %s nanoseconds to run!", name, elapsed)
}
