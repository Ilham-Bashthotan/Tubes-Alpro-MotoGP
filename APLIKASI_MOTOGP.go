package main

import (
	"fmt"
	"os"
	"os/exec"
)

/////////////////////////// Struct ///////////////////////////
const NMAX int = 24

type rider struct {
	no                                  int
	name, nat, team                     string
    FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC int
	poinSPR , poinRAC, totalPoin        int
}

type tabRider [NMAX]rider

var nDataRider int
var dataRider tabRider

/////////////////////////// Main Menu ///////////////////////////
func main() {
	var pilih string

	for {
		clear_screen()
		intro()
		menuUtama(&pilih)
		switch pilih {
			case "1": inputDataRider()
			case "2": tampilkanDataRider()
			case "3": editDataRider()
			case "4": hapusDataRider()
			default: clear_screen()
		} 
		if pilih == "5" {
			break
		}
	}
	bye()
}

/////////////////////////// Menu Cekat Tampilan ///////////////////////////
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
	fmt.Println("1.  Rider")
	fmt.Println("2.  Race Grid")
	fmt.Println("3.  Free Practice 1")
	fmt.Println("4.  Practice")
	fmt.Println("5.  Free Practice 2")
	fmt.Println("6.  Qualification 1")
	fmt.Println("7.  Qualification 2")
	fmt.Println("8.  Sprint Race")
	fmt.Println("9.  Warming Up")
	fmt.Println("10. Main Race")
	fmt.Println("11. Total Point")
	fmt.Println("12. Kembali")
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
	fmt.Println("1. Hapus Rider")
	fmt.Println("2. Hapus Semua Rider")
	fmt.Println("3. Kembali")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/3): ")
	inputFrasa(p)
}

/////////////////////////// Menu Tampilan ///////////////////////////
func inputDataRider() {
	var pilih string

	for {
		clear_screen()
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
			break
		}
	}
}

func tampilkanDataRider() {
	var pilih string

	for {
		clear_screen()
		menuTampilkanDataRider(&pilih)
		switch pilih {
			case "1": subMenuCetakRider()
			case "2": subMenuCetakRaceGrid()
			case "3": subMenuCetakFP1()
			case "4": subMenuCetakPR()
			case "5": subMenuCetakFP2()
			case "6": subMenuCetakQ1()
			case "7": subMenuCetakQ2()
			case "8": subMenuCetakSPR()
			case "9": subMenuCetakWUP()
			case "10": subMenuCetakRAC()
			case "11": subMenuCetakPoin()
			default: clear_screen()
		}
		if pilih == "12" {
			break
		}
	}
}

func editDataRider() {
	var pilih string

	for {
		clear_screen()
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
			break
		}
	}
}

func hapusDataRider() {
	var pilih string

	for {
		clear_screen()
		menuHapusDataRider(&pilih)
		switch pilih {
			case "1": 
			case "2":
			default: clear_screen()
		}
		if pilih == "3" {
			break
		}
	}
}

/////////////////////////// Sub Menu Tampilan  ///////////////////////////
func subMenuCetakRider() {
	var bin string
	
	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("        Data Rider        ")
	fmt.Println("--------------------------")
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakRider(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakRaceGrid() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("      Data Race Grid      ")
	fmt.Println("--------------------------")
	if nDataRider == 0 {
		fmt.Println("Data Rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !(cekKelengkapanQ1(dataRider, nDataRider) && cekKelengkapanQ2(dataRider, nDataRider)) {
		fmt.Println("Data Q1 atau Q2 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")		
	} else {
		cetakRider(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakPoin() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data Rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !(cekKelengkapanQ1(dataRider, nDataRider) && cekKelengkapanQ2(dataRider, nDataRider)) {
		fmt.Println("Data Q1 atau Q2 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakRider(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
	fmt.Print("Tekan enter sekali lagi. ")
}

func subMenuCetakFP1() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanFP1(dataRider, nDataRider) {
		fmt.Println("Data FP1 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuFP1(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
	fmt.Print("Tekan enter sekali lagi. ")
}

func subMenuCetakPR() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanPR(dataRider, nDataRider) {
		fmt.Println("Data PR masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuPR(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakFP2() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanFP2(dataRider, nDataRider) {
		fmt.Println("Data FP2 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuFP2(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakQ1() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanQ1(dataRider, nDataRider) {
		fmt.Println("Data Q1 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuQ1(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakQ2() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanQ2(dataRider, nDataRider) {
		fmt.Println("Data Q2 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuQ2(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakSPR() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanSPR(dataRider, nDataRider) {
		fmt.Println("Data SPR masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuSPR(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakWUP() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanWUP(dataRider, nDataRider) {
		fmt.Println("Data WUP masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuWUP(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

func subMenuCetakRAC() {
	var bin string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !cekKelengkapanRAC(dataRider, nDataRider) {
		fmt.Println("Data RAC masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakWaktuRAC(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	inputFrasa(&bin)
}

///////////////////////////  ///////////////////////////
func konversiWaktuTampilan(t, n int) string {
/*mengembalikan string dengan format “(menit)’(detik).(miliDetik)s” 
Jika posisi 1. Jika posisi > 1, maka mengembalikan string jarak dengan posisi pertama dengan. 
Jika waktu negatif, maka mengembalikan perolehan laps terakhir*/
	// var menit, detik, miliDetik int
	return ""
}

// func hitungTotalPoin(S rider) int {
// /*mengembalikan nilai total poin S rider*/
//     return S.poinSPR + S.poinRAC
// }

/////////////////////////// Urut Data ///////////////////////////
func urutNoRider(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu no*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.no < T[i - 1].no {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuFP1(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu FP1*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.FP1 < T[i - 1].FP1 {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuPR(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu PR*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.PR < T[i - 1].PR {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuFP2(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu FP2*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.FP2 < T[i - 1].FP2 {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuQ1(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu Q1*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.Q1 < T[i - 1].Q1 {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuQ2(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu Q2*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.Q2 < T[i - 1].Q2 {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuSPR(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu SPR*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.SPR < T[i - 1].SPR {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuWUP(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu WUP*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.WUP < T[i - 1].WUP {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutTotalPoin(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu totalPoin*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.totalPoin < T[i - 1].totalPoin {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuRAC(T *tabRider, n int) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu RAC*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && temp.RAC < T[i - 1].RAC {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func inputFrasa(str *string) {
/*I.S str 
F.S tercetak tabel dari field : name*/
	var ch byte

	*str = ""
	for ch != '\r'&& ch != '\n' {
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

/////////////////////////// Cetak Tabel ///////////////////////////
func cetakRider(T tabRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team*/
	var i int

	urutNoRider(&T, n)
	fmt.Printf("%25s %2s %3s %25s", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%20s %2.d %3s %25s", T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakRaceGrid(T tabRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team*/
	var i int

	// urut berdasarkan Q1 dan Q2

	// lakukan validasi bahwa data Q1 dan Q2 lengkap

	fmt.Printf("%25s %2s %3s %25s", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%20s %2.d %3s %25s", T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakWaktuFP1(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuFP1(&T, n)
	fmt.Printf("%25s %2s %3s %25s %11s", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%20s %2.d %3s %25s %11s", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP1, i))
	}
}

func cetakWaktuPR(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuPR(&T, n)
	fmt.Printf("%25v %2v %3v %25v %11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%20v %2.d %3v %25v %11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].PR, i))
	}
}

func cetakWaktuFP2(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuFP2(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP2, i))
	}
}

func cetakWaktuQ1(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuQ1(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q1, i))
	}
}

func cetakWaktuQ2(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuQ2(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q2, i))
	}
}

func cetakWaktuSPR(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuSPR(&T, n)
	fmt.Printf("%25v %2v %3v %25v %11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%20v %2v %3v %25v %11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].SPR, i))
	}
}

func cetakWaktuWUP(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuWUP(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].WUP, i))
	}
}

func cetakWaktuRAC(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuRAC(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-11v", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].RAC, i))
	}
}

func cetakPoin(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutTotalPoin(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-8v %-8v %-4v", "NAME", "NO", "NAT", "TEAM", "POIN SPR", "POIN RAC", "POIN")
	for i = 0; i < n; i++ {
		fmt.Printf("%-20v %-2v %-3v %-25v %-8v %-8v %-4v", T[i].name, T[i].no, T[i].nat, T[i].team, T[i].poinSPR, T[i].poinRAC, T[i].totalPoin)
	}
}

/////////////////////////// Edit ///////////////////////////


/////////////////////////// Hapus ///////////////////////////
func hapusData(T *tabRider, n *int, x int) {
	var i, idx int

	for i = idx; i < *n - 1; i++ {
		T[i] = T[i + 1]
		hapusBagian(T, i)
	} 
	*n--
}

func resetData(T *tabRider, n *int) {
	var i int

	for i = 0; i < NMAX; i++ {
		hapusBagian(T, i)
	} 
	*n = 0
}

func hapusBagian(T *tabRider, idx int) {
	T[idx].no = 0
	T[idx].FP1 = 0
	T[idx].PR = 0
	T[idx].FP2 = 0
	T[idx].Q1 = 0
	T[idx].Q2 = 0
	T[idx].SPR = 0
	T[idx].WUP = 0
	T[idx].RAC = 0
	T[idx].poinSPR = 0
	T[idx].poinRAC = 0
	T[idx].totalPoin = 0
	T[idx].name = ""
	T[idx].nat = ""
	T[idx].team = ""
}

/////////////////////////// Cek Kelengkapan Data ///////////////////////////
func cekKelengkapanRider(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP1(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanPR(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].PR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP2(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ1(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ2(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanSPR(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].SPR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanWUP(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].WUP == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanRAC(T tabRider, n int) bool {
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].RAC == 0 {
			return false
		}
	} 
	return true
}

/////////////////////////// Hiasan ///////////////////////////
func clear_screen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}