package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

var dataRider riders
var nDataRider int
var dataFP1, dataPR, dataFP2, dataQ1, dataQ2, dataSPR, dataWUP, dataRAC tabTime

func main() {
	var pilih string

	for {
		intro()
		menuUtama(&pilih)
		switch pilih {
		case "1": 
			inputDataRider()
		case "2": 
			tampilkanDataRider()
		case "3": 
			editDataRider()
		case "4": 
			hapusDataRider()
		default:
			clear_screen()
		} 
		if pilih == "5" {
			break
		}
	}
	bye()
}

// Tampilan
func intro() {
    clear_screen()
    fmt.Println("Selamat datang di aplikasi")
	fmt.Println("______  ___     _____      ________________")
	fmt.Println("___   |/  /_______  /________  ____/__  __ \\")
	fmt.Println("__  /|_/ /_  __ \\  __/  __ \\  / __ __  /_/ /")
	fmt.Println("_  /  / / / /_/ / /_ / /_/ / /_/ / _  ____/")
	fmt.Println("/_/  /_/  \\____/\\__/ \\____/\\____/  /_/")
}

func bye() {
    clear_screen()
    fmt.Println("Terima kasih telah menggunakan aplikasi MotoGP")
    fmt.Println("By:")
    fmt.Println("Ilham Bashthotan")
    fmt.Println("Faiz Nawfal Zulfiqar")
}

func menuUtama(p *string) {
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Input Data Rider")
	fmt.Println("2. Tampilkan Data Rider")
	fmt.Println("3. Edit Data Rider")
	fmt.Println("4. Hapus Data Rider")
	fmt.Println("5. Keluar")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1-5): ")
	inputFrasa(p)
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

func menuTampilkanDataRider(p *string) {
	fmt.Println("--------------------------")
	fmt.Println("   Tampilkan Data Rider   ")
	fmt.Println("--------------------------")
	fmt.Println("1.  Race Grid")
	fmt.Println("2.  Free Practice 1")
	fmt.Println("3.  Practice")
	fmt.Println("4.  Free Practice 2")
	fmt.Println("5.  Qualification 1")
	fmt.Println("6.  Qualification 2")
	fmt.Println("7.  Sprint Race")
	fmt.Println("8.  Warming Up")
	fmt.Println("9.  Main Race")
	fmt.Println("10. Total Point")
	fmt.Println("11. Kembali")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1-11): ")
	inputFrasa(p)
}

func menuEditDataRider(p *string) {
	fmt.Println("--------------------------")
	fmt.Println("     Edit Data Rider      ")
	fmt.Println("--------------------------")
	fmt.Println("1.  Rider")
	fmt.Println("2.  Free Practice 1")
	fmt.Println("3.  Practice")
	fmt.Println("4.  Free Practice 2")
	fmt.Println("5.  Qualification 1")
	fmt.Println("6.  Qualification 2")
	fmt.Println("7.  Sprint Race")
	fmt.Println("8.  Warming Up")
	fmt.Println("9.  Main Race")
	fmt.Println("10. Kembali")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1-10): ")
	inputFrasa(p)
}

func menuHapusDataRider(p *string) {
	fmt.Println("--------------------------")
	fmt.Println("     Hapus Data Rider     ")
	fmt.Println("--------------------------")
	fmt.Println("1. Rider")
	fmt.Println("2. Semua Rider")
	fmt.Println("3. Kembali")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/3): ")
	inputFrasa(p)
}

func inputDataRider() {
	var pilih string

	clear_screen()
	for {
		menuInputDataRider(&pilih)
		switch pilih {
			// case "1":	inputRider(&dataRider, nDataRider)
			// case "2":	inputFP1(&dataRider, nDataRider)
			// case "3":	inputPR(&dataRider, nDataRider)
			// case "4":	inputFP2(&dataRider, nDataRider)
			// case "5":	inputQ1(&dataRider, nDataRider)
			// case "6":	inputQ2(&dataRider, nDataRider)
			// case "7":	inputSPR(&dataRider, nDataRider)
			// case "8":	inputWUP(&dataRider, nDataRider)
			// case "9":	inputRAC(&dataRider, nDataRider)
			default: clear_screen()
		}
		if pilih == "10" {
			clear_screen()
			break
		}
	}
}

func tampilkanDataRider() {
	var pilih string

	clear_screen()
	for {
		menuTampilkanDataRider(&pilih)
		switch pilih {
			// case "1":	cetakTabelPemain(dataRider, nDataRider)
			// case "2":	cetakTabelWaktu(dataFP1, nDataRider)
			// case "3":	cetakTabelWaktu(dataPR, nDataRider)
			// case "4":	cetakTabelWaktu(dataFP2, nDataRider)
			// case "5":	cetakTabelWaktu(dataQ1, nDataRider)
			// case "6":	cetakTabelWaktu(dataQ2, nDataRider)
			// case "7":	cetakTabelWaktu(dataSPR, nDataRider)
			// case "8":	cetakTabelWaktu(dataWUP, nDataRider)
			// case "9":	cetakTabelWaktu(dataRAC, nDataRider)
			// case "10": cetakTabelPoin(dataRider, nDataRider)
			default: clear_screen()
		}
		if pilih == "11" {
			clear_screen()
			break
		}
	}
}

func editDataRider() {
	var pilih string

	clear_screen()
	for {
		menuEditDataRider(&pilih)
		switch pilih {
			case "1": 
			case "2":
			case "3":
			case "4":
			case "5":
			case "6":
			case "7":
			case "8":
			case "9":
			default: clear_screen()
		}
		if pilih == "10" {
			clear_screen()
			break
		}
	}
}

func hapusDataRider() {
	var pilih string
/*
	fmt.Println("1. Rider")
	fmt.Println("2. Semua Rider")
	fmt.Println("3. Kembali")
*/
	clear_screen()
	for {
		menuHapusDataRider(&pilih)
		switch pilih {
			case "1": 
			case "2":
			default: clear_screen()
		}
		if pilih == "3" {
			clear_screen()
			break
		}
	}
}

// func konversiWaktuTampilan(T  tabSubRider, posisi  integer) string {
// /*mengembalikan string dengan format “(menit)’(detik).(miliDetik)s” jika posisi 1. Jika posisi > 1, maka mengembalikan string jarak dengan posisi pertama dengan . Jika waktu negatif, maka mengembalikan perolehan laps terakhir*/
// 	var menit, detik, miliDetik int

// 	T[]
// 	return
// }

// func hitungTotalPoin(S rider) int {
// /*mengembalikan nilai total poin S rider*/
//     return S.poinSPR + S.poinRAC
// }



// func inputRider(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai. Banyak elemen (n) berisi nilai NMAX, jika n > NMAX*/
// } 

// func inputFP1(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeFP1*/
// }

// func inputPR(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timePR*/
// }

// func inputFP2(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeFP2*/
// }

// func inputQ1(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeQ1*/
// }

// func inputQ2(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeQ2*/
// }

// func inputSPR(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeSPR*/
// }

// func inputWUP(T *tabRider, n int) {
// 	/*I.S. terdefinisi array T sebanyak n
// 	F.S. T berisi nilai memasukkan field timeWUP*/
// 	}

// func inputRAC(T *tabRider, n int) {
// /*I.S. terdefinisi array T sebanyak n
// F.S. T berisi nilai memasukkan field timeRAC*/
// }

// func urutWaktu(T *tabSubRider) {
// /*I.S. terdefinisi array acak T1 sebanyak n
// F.S. T berisi nilai terurut berdasarkan waktu*/
// 	var i, minPos, pass int
// 	var temp subRider

// 	for pass = 1; pass < 6; pass++ {
// 		minPos = pass - 1
// 		for i = pass; i < 6; i++ {
// 			if T[i].time < T[minPos].time {
// 				minPos = i
// 			}
// 		}
// 		temp = T[minPos]
// 		T[minPos] = T[pass - 1]
// 		T[pass - 1] = temp
// 	}
// }

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

// Hiasan
func clear_screen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

// func cetakTabelPemain(T tabRider, n int) {
// /*I.S. T array dan banyak elemen n terdefinisi
// F.S. tercetak tabel dari field : name, no, nat, team*/

// }

// func cetakTabelWaktu(T tabSubRider, n int) {
// /*I.S. T array dan banyak elemen n terdefinisi
// F.S. tercetak tabel dari field : name, no, nat, team, time*/

// }

// func cetakTabelPoin(T tabRider, n int) {
// /*I.S T array dan banyak elemen n terdefinisi
// 	F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
// }