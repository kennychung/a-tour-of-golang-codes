package main

import "fmt"

type Vertex struct{
    X int
    Y int
}

func main() {
    p := Vertex{1,2}
    q := &p
    q.X = 1e2
    fmt.Println(p) //Go has pointers, but no pointer arithmetic.
}
