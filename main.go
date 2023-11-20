package main

import (
    "fmt"
    "os"
    "cryptomata/feistel"
)

func main() {
    if len(os.Args) < 3 {
        abort("a minimum of two argument are required.")
    }
    input := os.Args[1:]
    output := feistel.Feistel(input[:])
    fmt.Println(output)
}

func abort(s string) {
    fmt.Println(s)
    os.Exit(1)
}
