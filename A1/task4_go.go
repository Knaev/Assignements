#!/usr/bin/go
import (
    . "fmt" // Using '.' to avoid prefixing functions with their package names
    . "runtime" // This is probably not a good idea for large projects...
    . "time"
)

var i = 0

func adder() {
    for x := 0; x < 1000000; x++ {
        i++
    }
}
func subber() {
    for x := 0; x < 1000000; x++ {
        i--
    }
}

func main() {
    GOMAXPROCS(NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
    go adder() // This spawns adder() as a goroutine
    go subber() // This spawns subber() as a goroutine
    Sleep(100*Millisecond)
    Println("Done:", i);
}
