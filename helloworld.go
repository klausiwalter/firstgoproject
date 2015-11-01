package main

import (
    "fmt"
)

func main() {

    a, b := swap(1, 2)

    fmt.Println(a, b)

}

func swap(x, y int) (int, int) {

    return y, x
}
