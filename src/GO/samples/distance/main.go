package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("input text:")
    var n1, n2 float32
    n, err := fmt.Scanln(&n1, &n2)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("number of items read: %d\n", n)
    fmt.Printf("read line: %s %s\n", n1, n2)
}