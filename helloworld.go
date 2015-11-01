package main

import (
    "fmt"
)

func main() {

    fmt.Println(increment(1))

}

func increment(input int) (output int) {

    output = input + 1

    return
}
