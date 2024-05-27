package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
)

type mahasiswa struct {
    nama string
    nim int
}

const NMAX int = 10

type tabMhs [NMAX]mahasiswa

func main() {
    var kata string
    var data tabMhs
    var ndata int

    // fmt.Scan(&kata)
    // fmt.Println(kata)
    // fmt.Println(len(kata))

    fmt.Println("______  ___     _____      ________________")
    fmt.Println("___   |/  /_______  /________  ____/__  __ \\")
    fmt.Println("__  /|_/ /_  __ \\  __/  __ \\  / __ __  /_/ /")
    fmt.Println("_  /  / / / /_/ / /_ / /_/ / /_/ / _  ____/")
    fmt.Println("/_/  /_/  \\____/\\__/ \\____/\\____/  /_/")
    
    fmt.Printf("+-%s-+\n", ulang("-", len(kata)))
    fmt.Printf("| %v |\n",kata)
    fmt.Printf("+-%s-+\n", ulang("-", len(kata)))
    // clear_screen()
    fmt.Scan(&ndata)
    inputData(&data, &ndata)
    outputData(data, ndata)
    // clear_screen()
}

func inputData(A *tabMhs, n *int) {
    var i int 

    if *n > NMAX {
        *n = NMAX
    }
    for i = 0; i <= *n-1; i++ {
        fmt.Print("MAUKKAN NAMA: ")
        inputFrasa(&A[i].nama)
        fmt.Print("MAUKKAN NIM: ")
        fmt.Scan(&A[i].nim)
    }
}
/* 
8
dwi saputra
53152531 
badru lasa
45132131325253 
agus salim
14541613 
bina raga
14525251 
asep jamil
14531346 
buyu 
14513514
cantika saputri
14613514
geling udu
23513513
*/ 
// Hanya mempercantik tampilan
func outputData(A tabMhs, n int) {
    var i, maxNama int 
    maxNama = maxRows(A, n, "nama")
    maxNim := maxRows(A, n, "nim")
    fmt.Printf(".%s...%s.\n",  ulang(".", maxNama), ulang(".", maxNim))
    fmt.Printf(" %s%s   %s%s\n", "NAMA", ulang(" ", maxNama - 4), "NIM", ulang(" ", maxNim - 3))
    fmt.Printf(".%s...%s.\n",  ulang(".", maxNama), ulang(".", maxNim))
    for i = 0; i <= n-1; i++ {
        fmt.Printf(" %v%s   %d%s  \n", A[i].nama, ulang(" ", maxNama - len(A[i].nama)), A[i].nim, ulang(" ", maxNim - len(strconv.Itoa(A[i].nim))))
    }
    fmt.Printf(".%s...%s.\n",  ulang(".", maxNama), ulang(".", maxNim))
}

func maxRows(A tabMhs, n int, fieldName string) int {
    var max int

    for i := 0; i < n; i++ {
        v := reflect.ValueOf(A[i])
        field := v.FieldByName(fieldName)

        var length int
        switch field.Kind() {
        case reflect.String:
            length = len(field.String())
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            length = len(strconv.FormatInt(field.Int(), 10))
        }

        if length > max {
            max = length
        }
    }

    return max
}
func ulang(str string, n int) string {
    var i int
    var hasil string

    hasil = ""
    for i = 1; i <= n; i++ {
        hasil += str
    }
    return hasil
}


func inputFrasa(str *string) {
    var ch, i byte

    *str = ""
    for i = 1; i <= 2; i++ {
        fmt.Scanf("%c", &ch)
        if ch != '\r' && ch != '\n' {
            *str += string(ch)
        }
    }
    for {
        fmt.Scanf("%c", &ch)
        if ch != '\r' && ch != '\n' {
            *str += string(ch)
        }
        if ch == '\n' {
            break
        }
    }
}

func clear_screen() {
    c := exec.Command("cmd", "/c", "cls")
    c.Stdout = os.Stdout
    c.Run()
}

/*
+2.280s
+4.174s
+4.798s
+7.698s
+9.185s
+11.190s
+11.516s
+12.257s
+12.699s
+13.492s
+16.439s
+16.816s
+16.969s
+19.123s
+23.618s
+27.854s
*/






