package main

import (
  "fmt"
)

// given a starting number N
// if the nunber is even, the next nunber is N / 2
// if the number is odd, the next number is 3 * N + 1
// print all numbers from N until 1

//recursive unoptimized hailStone sequence
func hailStone( num uint32, sequence []uint32 ) []uint32 {
  sequence = append( sequence, num )

  if ( num == 1 ) {
    return sequence
  }

  if ( num % 2 == 0 ) {
    sequence = hailStone( num / 2, sequence )
  } else {
    sequence = hailStone( 3 * num + 1, sequence )
  }

  return sequence
}

func printHailStone( arr []uint32 ) {
  for _, value := range arr {
    fmt.Printf("%d->", value)
  }
  fmt.Printf("done\n")
}

func main() {
  var i uint32 = 1
  for ; i < 100; i++ {
    var arr []uint32

    printHailStone(hailStone( i, arr ) )
  }
}
