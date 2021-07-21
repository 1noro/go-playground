package main

import (
    "os"
    "fmt"
    //"bufio"
    "strconv"
)

// ----------------------------------------------------------------------------
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// ----------------------------------------------------------------------------
func printMenu() {
    fmt.Println("--- MENU ---")
    fmt.Println("1.- Altas")
    fmt.Println("2.- Visualizar fichero")
    fmt.Println("3.- Crear ficheros por sexos")
    fmt.Println("4.- Fin")
}

func readOpt() int {
    out := 0
    for out < 1 || out > 4 {
        fmt.Print("Escribe: ")
        fmt.Scanf("%d", &out)
    }
    return out
}

// ----------------------------------------------------------------------------
func altas(filename string) {
    fmt.Println("\nAltas\n")
    // var nombre string
    // var edad int
    // var sexo string
    f, err := os.Create(filename)
    check(err)
    _, err = f.WriteString("Hello World\n")
    check(err)
    err = f.Close()
    check(err)
}

// ----------------------------------------------------------------------------
func main() {
    /*var filnemae = "personas.txt"
    var opt int*/

    i, err := strconv.Atoi("-42")
    check(err)
    fmt.Println(i)

    s := strconv.Itoa(-42)
    fmt.Println(s)

    /*for opt != 4 {
        printMenu()
        opt = readOpt()
        switch opt {
        case 1:
            altas(filnemae)
        case 2:
            fmt.Println("\nListados\n")
        case 3:
            fmt.Println("\nSeparar\n")
        }
    }

    fmt.Println("Bye :(")*/
}
