package main

import (
    "crypto/sha1"
    "fmt"
    "io"
    "encoding/hex"
)

func main() {
    calculator1 := sha1.New()
	calculator2 := sha1.New()
	
    io.WriteString(calculator1, "Alfredo")
    io.WriteString(calculator2, "Alfredoo")

    hashByteArray1 := calculator1.Sum(nil)
    hashByteArray2 := calculator2.Sum(nil)

    // fmt.Printf("%x\n", hashByteArray)
    // fmt.Printf("%X\n", hashByteArray)
    // fmt.Printf("% x\n", hashByteArray)
    // fmt.Printf("% X\n", hashByteArray)
    // fmt.Println(hashByteArray)

    // str := hex.EncodeToString(hashByteArray)
	// fmt.Println(str)
	
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
