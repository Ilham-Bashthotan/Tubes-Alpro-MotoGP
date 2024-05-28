package main

import "fmt"

const NMAX int = 24

type time struct {
	no, time int
}

type tabTime [NMAX]time

type riders struct {
	no                                  [NMAX]int
	name, nat, team                     [NMAX]string
    FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC tabTime
}

func main() {
	var dataRider riders
	var nRider int

	inputDataRider(&dataRider, &nRider)
	cetakDataRider(dataRider, nRider)
}

func inputDataRider(T *riders, n *int) {
	var i int

	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i = 0; i < *n; i++ {
		fmt.Println("Input nama rider:")
		inputFrasa(&T.name[i])
		fmt.Println("Input nomor rider:")
		fmt.Scan(&T.no[i])
		fmt.Println("Input asal negara rider:")
		fmt.Scan(&T.nat[i])
		fmt.Println("Input tim rider:")
		fmt.Scan(&T.team[i])
		fmt.Println("Input nomor & waktu FP1:")
		fmt.Scan(&T.FP1[i].no, &T.FP1[i].time)
		fmt.Println("Input nomor & waktu PR:")
		fmt.Scan(&T.PR[i].no,&T.PR[i].time)
		fmt.Println("Input nomor & waktu FP2:")
		fmt.Scan(&T.FP2[i].no, &T.FP2[i].time)
		fmt.Println("Input nomor & waktu Q1:")
		fmt.Scan(&T.Q1[i].no, &T.Q1[i].time)
		fmt.Println("Input nomor & waktu Q2:")
		fmt.Scan(&T.Q2[i].no, &T.Q2[i].time)
		fmt.Println("Input nomor & waktu SPR:")
		fmt.Scan(&T.SPR[i].no, &T.SPR[i].time)
		fmt.Println("Input nomor & waktu WUP:")
		fmt.Scan( &T.WUP[i].no, &T.WUP[i].time)
		fmt.Println("Input nomor & waktu RAC:")
		fmt.Scan( &T.RAC[i].no, &T.RAC[i].time)
	}
}

func cetakDataRider(T riders, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(T.name[i], T.no[i], T.nat[i], T.team[i], T.FP1[i].no, T.FP1[i].time, T.PR[i].no,T.PR[i].time, T.FP2[i].no, T.FP2[i].time, T.Q1[i].no, T.Q1[i].time, T.Q2[i].no, T.Q2[i].time, T.SPR[i].no, T.SPR[i].time, T.WUP[i].no, T.WUP[i].time, T.RAC[i].no, T.RAC[i].time)
	}
}

func inputFrasa(str *string) {
	/*I.S str 
	F.S tercetak tabel dari field : name*/
		var ch, i byte
	
		*str = ""
		for i = 1; i <= 2; i++ {
			fmt.Scanf("%c", &ch)
			if ch != '\r'&& ch != '\n' {
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