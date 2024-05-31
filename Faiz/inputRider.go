package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX int = 24

type rider struct {
	no, FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC   int
	name, nat, team                           string
}

type tabRider [NMAX]rider

func main() {
	var dataRider tabRider
	var nRider int
	var pilih string

	for {
		menuInputDataRider(&pilih)
		switch pilih {
			case "1": inputRider(&dataRider, &nRider, "Rider")
			case "2": inputRider(&dataRider, &nRider, "FP1")
			case "3": inputRider(&dataRider, &nRider, "PR")
			case "4": inputRider(&dataRider, &nRider, "FP2")
			case "5": inputRider(&dataRider, &nRider, "Q1")
			case "6": inputRider(&dataRider, &nRider, "Q2")
			case "7": inputRider(&dataRider, &nRider, "SPR")
			case "8": inputRider(&dataRider, &nRider, "WUP")
			case "9": inputRider(&dataRider, &nRider, "RAC")
		    default: clear_screen()
 		}
		if pilih == "10" {
			clear_screen()
			break
		}
	}
	cetakDataRider(dataRider, nRider)
}

func inputRider(T *tabRider, n *int, p string) {
	var i, min, sec, milsec int
	// var bin string

	if p == "Rider"{
		fmt.Println("Masukkan jumlah rider")
		fmt.Scan(n)
		if *n > NMAX {
			*n = NMAX
		}
		for i = 0; i < *n; i++ {
			fmt.Println("Input nama rider ke", i+1, ":")
			inputFrasa(&T[i].name)
			fmt.Println("Input nomor rider ke", i+1, ":")
			fmt.Scan(&T[i].no)
			fmt.Println("Input asal negara rider ke", i+1, ":")
			inputFrasa(&T[i].nat)
			fmt.Println("Input tim rider ke", i+1, ":")
			inputFrasa(&T[i].team)
		}
	} else if p == "FP1" {
		if cekKelengkapanRider(*T, *n){
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu FP1 rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].FP1 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "PR" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu PR rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].PR = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "FP2" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu FP2 rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].FP2 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "Q1" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu Q1 rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].Q1 = (60000 * min) + (1000 * sec) + milsec
			}
		}	
	} else if p == "Q2" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu Q2 rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].Q2 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "SPR" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu SPR rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].SPR = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "WUP" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu WUP rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].WUP = (60000 * min) + (1000 * sec) + milsec
			}
		}
		
	} else if p == "RAC" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu RAC rider ke", i+1, "dalam format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].RAC = (60000 * min) + (1000 * sec) + milsec 
			}
		}
	}
}

func cetakDataRider(T tabRider, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println(T[i].name, T[i].no, T[i].nat, T[i].team, T[i].FP1, T[i].PR, T[i].PR, T[i].FP2, T[i].Q1, T[i].Q2, T[i].SPR, T[i].WUP, T[i].RAC)
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

	func menuInputDataRider(p *string) {
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
		inputFrasa(p)
	}

	func clear_screen() {
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	}

	func cekKelengkapanRider(T tabRider, n int) bool {
		var i int
		
		for i = 0; i < n; i++ {
			if T[i].no == 0 {
				return false
			}
		} 
		return true
	}