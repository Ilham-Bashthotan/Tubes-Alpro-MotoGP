package main

import "fmt"

type waktu struct {
	detik, no int
}

type tabWaktu [NMAX]waktu

type rider struct {
	no [NMAX]int
	nama, nat [NMAX]string
	FP1, PR, FP2 tabWaktu
}

const NMAX int = 24


// type tabRider [NMAX]rider

func main() {
	var dataRider rider
	var nDataRider int
	var waktuFP1, waktuPR, waktuFP2 tabWaktu 

	fmt.Scan(&nDataRider)
	// input rider
	for i := 0; i < nDataRider; i++ {
		fmt.Scan(&dataRider.no[i],&dataRider.nama[i],&dataRider.nat[i])
		dataRider.FP1[i].no = dataRider.no[i]
		dataRider.PR[i].no = dataRider.no[i]
		dataRider.FP2[i].no = dataRider.no[i]
	}

	// input waktu FP1
	fmt.Println("input FP1:")
	for i := 0; i < nDataRider; i++ {
		fmt.Println(dataRider.nama[i])
		fmt.Scan(&dataRider.FP1[i].detik)
	}

	// input waktu PR
	fmt.Println("input PR:")
	for i := 0; i < nDataRider; i++ {
		fmt.Println(dataRider.nama[i])
		fmt.Scan(&dataRider.PR[i].detik)
	}

	// input waktu FP2
	fmt.Println("input PR:")
	for i := 0; i < nDataRider; i++ {
		fmt.Println(dataRider.nama[i])
		fmt.Scan(&dataRider.FP2[i].detik)
	}

	// ngurutin waktu
	waktuFP1 = urut(dataRider.FP1, nDataRider)
	waktuPR = urut(dataRider.PR, nDataRider)
	waktuFP2 = urut(dataRider.FP2, nDataRider)

	fmt.Println(waktuFP1)
	fmt.Println(waktuPR)
	fmt.Println(waktuFP2)
}

func urut(T tabWaktu, n int) tabWaktu {
	// utut membesar
    var i, minPos, pass int
    var temp waktu
    
    for pass = 1; pass < n; pass++ {
        minPos = pass - 1
        for i = pass; i < n; i++ {
            if T[i].detik < T[minPos].detik {
                minPos = i
            }
        }
        temp = T[minPos]
        T[minPos] = T[pass - 1]
        T[pass - 1] = temp
    }
	return T
}
/*
2
13 integralTakWajar MRI
10 arcimedes FIS

//FP1
123
53223

//PR
1234432
13413 

//FP2
234
235235
*/ 