package main

import(
    "fmt"
)

func main() {
    fmt.Println(split(17))
}

func split(sum int)(x,y int){
    x = sum * 2
    y = x - sum
    return
}
