package main

import "fmt"

func swap(x,y string) (string,string){
    return y,x
}

func main() {
    a,b := swap("first","last")
    fmt.Println(a,b)
    c,d := swap(a,b)
    fmt.Println(c,d)
}
