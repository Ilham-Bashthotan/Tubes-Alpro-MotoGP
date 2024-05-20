package main

import "fmt"

type pemain struct {
    nama string
    FP1, PR, FP2, Q1, Q2 int
}

type subPemain struct {
    nama string
    waktu int
}

type tabPemain [6]pemain

type tabSubPemain [6]subPemain

func main() {
    var data tabPemain
    var data1, data2, data3 tabSubPemain

    inputData(&data)
    // memindahkan data
    for i := 0; i < 6; i++ {
        data1[i].nama = data[i].nama
        data1[i].waktu = data[i].FP1
    }

    for i := 0; i < 6; i++ {
        data2[i].nama = data[i].nama
        data2[i].waktu = data[i].PR
    }

    for i := 0; i < 6; i++ {
        data3[i].nama = data[i].nama
        data3[i].waktu = data[i].FP2
    }

    urut(&data1)
    urut(&data2)
    urut(&data3)

    cetak(data1)
    cetak(data2)
    cetak(data3)
    
}
func inputData(T *tabPemain) {
    for i := 0; i < 6; i++ {
        fmt.Scan(&T[i].nama, &T[i].FP1, &T[i].PR, &T[i].FP2)
    }
}

func urut(T *tabSubPemain) {
    var i, minPos, pass int
    var temp subPemain
    
    for pass = 1; pass < 6; pass++ {
        minPos = pass - 1
        for i = pass; i < 6; i++ {
            if T[i].waktu < T[minPos].waktu {
                minPos = i
            }
        }
        temp = T[minPos]
        T[minPos] = T[pass - 1]
        T[pass - 1] = temp
    }
}


func cetak(T tabSubPemain) {
    fmt.Println("\nData adalah:")
    for i := 0; i < 6; i++ {
        fmt.Println(T[i].nama, T[i].waktu)
    }
}
/*
a 123 1212 432
b 213 2144 4
c 34 414 41
d 4141 42 3
e 428 272 26
f 26 67 93
*/
