package main

import (
    "fmt"
    "math/rand"
)

func main() {

    fmt.Println(generateRandomNumber(100))

}

func generateRandomNumber(max int) int {

    return rand.Intn(max)
}
