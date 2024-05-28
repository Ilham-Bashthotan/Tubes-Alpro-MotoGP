package main

import "fmt"

const NMAX int = 24

type rider struct {
	no                                  int
	name, nat, team                     string
    FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC int //milidetik
}

type tabRider [NMAX]rider

func main() {
	var dataRider riders
	var nRider int
	var pilih string

	for {
		fmt.Scan(&pilih)
		switch pilih {
			case "1": inputRider(&dataRider, &nRider, "rider")
			case "2": inputRider(&dataRider, &nRider, "FP1")
			case "3": inputRider(&dataRider, &nRider, "PR")
			case "4": inputRider(&dataRider, &nRider, "FP2")
			case "5": inputRider(&dataRider, &nRider, "Q1")
			case "6": inputRider(&dataRider, *nRider, "Q2")
			case "7": inputRider(&dataRider, *nRider, "SPR")
			case "8": inputRider(&dataRider, *nRider, "WUP")
			case "9": inputRider(&dataRider, *nRider, "RAC")
 		}
	}
	cetakDataRider(dataRider, nRider)
}

func inputRider(T *tabRider, n *int, p string) {
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

func cetakDataRider(T tabRider, n int) {
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

	func menuInputDataRider() {
		fmt.Println("--------------------------")
		fmt.Println("     Input Data Rider     ")
		fmt.Println("--------------------------")
		fmt.Println("1.  Rider")
		fmt.Println("2.  Waktu Free Practice 1")
		fmt.Println("3.  Waktu Practice")
		fmt.Println("4.  Waktu Free Practice 2")
		fmt.Println("5.  Waktu Qualification 1")
		fmt.Println("6.  Waktu Qualification 2")
		fmt.Println("7.  Waktu Sprint Race")
		fmt.Println("8.  Waktu Warming Up")
		fmt.Println("9.  Waktu Main Race")
		fmt.Println("10. Kembali")
		fmt.Println("--------------------------")
		fmt.Print("Pilih (1-10): ")
	}

