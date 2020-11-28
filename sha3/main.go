package main

import (
    "golang.org/x/crypto/sha3" // TO INSTALL RUN: go get -u golang.org/x/crypto/...
    "fmt"
    "io"
    "encoding/hex"
)

func main() {
    calculator1 := sha3.New224()
    calculator2 := sha3.New224()
    
    io.WriteString(calculator1, "Alfredo")
    io.WriteString(calculator2, "Alfredoo")

    hashByteArray1 := calculator1.Sum(nil)
    hashByteArray2 := calculator2.Sum(nil)

    str1 := hex.EncodeToString(hashByteArray1)
    str2 := hex.EncodeToString(hashByteArray2)

    fmt.Println(str1)
    fmt.Println(str2)
    
    for i, char1 := range str1 {
        if string(char1) == string(str2[i]) {
            fmt.Print("^")
        } else {
            fmt.Print(" ")
        }
    }
    fmt.Print("\n")
}
