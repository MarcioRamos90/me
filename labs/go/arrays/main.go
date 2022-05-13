package main

import (
    "fmt"
)

func main() {

    var arr [5]int
    fmt.Println("1: ", arr)

    arr[2] = 342
    fmt.Println("2: ", arr)

    fmt.Println("3: len", len(arr))

    arr2 := [5]int{1,2,3,4,5}
    fmt.Println("arr2:", arr2)

    var arr3 [2][2]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            arr3[i][j] = i + j
        }
    }
    fmt.Println("2d: ", arr3)
}
