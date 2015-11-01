package main

import (
    "fmt"
)

func main() {

    var lowValue, highValue int = 1, 1000

    fmt.Println(increment(lowValue))
    fmt.Println(increment(highValue))

}

func increment(input int) (output int) {

    output = input + 1

    return
}
