package main

import (
    "systats"
    "fmt"
)

func main(){
    for k,v := range systats.GetDiskFree() {
        fmt.Printf("%s -> %.5f\n", k, v)
    }
}
