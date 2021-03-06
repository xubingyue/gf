package main

import (
    "fmt"
    "gitee.com/johng/gf/g/os/gtime"
)

func main() {
    fmt.Println("Date       :", gtime.Date())
    fmt.Println("Datetime   :", gtime.Datetime())
    fmt.Println("Second     :", gtime.Second())
    fmt.Println("Millisecond:", gtime.Millisecond())
    fmt.Println("Microsecond:", gtime.Microsecond())
    fmt.Println("Nanosecond :", gtime.Nanosecond())
}
