package main

import (
  "fmt"
)

// Challenge
// iterate through the numbers 1 to N
// if it is divisible by 3, print Fizz
// if it is diviusble by 5, print Buzz
// if it not divisible 3 or 5, print the number



func FizzBuzz( maxNum uint64 ) {
  var num uint64
  for num = 1; num <= maxNum; num++ {

    if ( num % 3 == 0 ) {
      fmt.Printf("Fizz")
    }
    if ( num % 5 == 0) {
      fmt.Printf("Buzz")
    }

    if ( !( num % 3 == 0 || num % 5 == 0 ) ) {
      fmt.Printf("%d", num)
    }

    fmt.Printf("\n")
  }

}

func FizzBuzz2( maxNum uint64 ) {
  var index uint64

   for index = 1; index < maxNum; index++ {

    if ( index % 3 == 0 ) {
      fmt.Printf("Fizz")
    }

    if ( index % 5 == 0 ) {
      fmt.Printf("Buzz")
    }

    if ( !( index % 3 == 0 || index % 5 == 0)  ) {
      fmt.Printf("%d", index)
    }

    fmt.Println()

  }

}


func main() {
  FizzBuzz2(64)
}
