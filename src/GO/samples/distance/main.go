package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {
    var r_distance float64
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    n1 , err := strconv.ParseFloat(os.Args[1],32)
    n2 , err := strconv.ParseFloat(os.Args[2],32)
    r_distance = n2/n1*100-100
    fmt.Println(err)
    fmt.Printf("read line: %s %s\n", n1, n2)
    fmt.Printf("result: %.2f%%\n",r_distance)
    fmt.Println(argsWithProg)
    fmt.Println("---")
    fmt.Println(argsWithoutProg)
}