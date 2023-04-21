package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // [Vulnerability] Hardcoded password
    password := []byte("Alice and Bob...")

    /*
     * Go 1.17 added conversions from slice to an array pointer. Go 1.20 extends this to allow conversions
     * from a slice to an array: given a slice x, [4]byte(x) can now be written instead of *(*[4]byte)(x).
    */
    xArray := [16]byte(password)  // illegal in Go < 1.20
    xSlice := xArray[:]

    // [Vulnerability] Privacy Violation (the password is being sent to the console)
    fmt.Printf("%s\n", xSlice)

    /*
     * unsafe.String is a new 1.20 function in the "unsafe" package of Go
     */
    bytePtr := &xArray[0]
    xString := unsafe.String(bytePtr, 16)
    // [Vulnerability] Privacy Violation (the password is being sent to the console)
    fmt.Printf("%s\n", xString)
}
