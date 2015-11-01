package main

import (
    "fmt"
)

func main() {

    var lowValue int = 1

    highValue := 1000

    fmt.Println(increment(lowValue))
    fmt.Println(increment(highValue))

    var zeroInteger int
    var zeroString string
    var zeroBoolean bool

    fmt.Println(zeroString)
    fmt.Println(zeroInteger)
    fmt.Println(zeroBoolean)

}

func increment(input int) (output int) {

    output = input + 1

    return
}
