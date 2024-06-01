package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

/////////////////////////// Struct ///////////////////////////
const NMAX int = 24

type rider struct {
	no                                  int
	name, nat, team                     string
    FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC int
	poinSPR, poinRAC, totalPoin         int
}

type tabRider [NMAX]rider

var nDataRider int
var dataRider tabRider

/////////////////////////// Main Menu ///////////////////////////
func main() {
	var pilih string
	// test(&dataRider, &nDataRider)
	// hitungTotalPoin(&dataRider, nDataRider)
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

/////////////////////////// Cekat Main Menu ///////////////////////////
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
			case "1": inputRider(&dataRider, &nDataRider, pilih)
			case "2": inputRider(&dataRider, &nDataRider, pilih)
			case "3": inputRider(&dataRider, &nDataRider, pilih)
			case "4": inputRider(&dataRider, &nDataRider, pilih)
			case "5": inputRider(&dataRider, &nDataRider, pilih)
			case "6": inputRider(&dataRider, &nDataRider, pilih)
			case "7": inputRider(&dataRider, &nDataRider, pilih)
			case "8": inputRider(&dataRider, &nDataRider, pilih)
			case "9": inputRider(&dataRider, &nDataRider, pilih)
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
			case "1": subMenuHapusRider()
			case "2": subMenuHapusSemuaRider()
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
	fmt.Scanln(&bin)
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
	fmt.Scanln(&bin)
}

func subMenuCetakPoin() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("        Data Poin         ")
	fmt.Println("--------------------------")
	if nDataRider == 0 {
		fmt.Println("Data Rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else if !(cekKelengkapanQ1(dataRider, nDataRider) && cekKelengkapanQ2(dataRider, nDataRider)) {
		fmt.Println("Data Q1 atau Q2 masih belum lengkap.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		cetakPoin(dataRider, nDataRider)
	}
	fmt.Print("Tekan enter untuk keluar. ")
	fmt.Scanln(&bin)
}

func subMenuCetakFP1() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data FP1         ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
	fmt.Print("Tekan enter sekali lagi. ")
}

func subMenuCetakPR() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data PR          ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakFP2() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data FP2         ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakQ1() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data Q1          ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakQ2() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data Q2          ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakSPR() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data SPR         ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakWUP() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data WUP         ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuCetakRAC() {
	var bin string

	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("         Data RAC         ")
	fmt.Println("--------------------------")
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
	fmt.Scanln(&bin)
}

func subMenuHapusRider() {
	var pilih string
	var x, idx int

	for {
		clear_screen()
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			for {
				cetakRider(dataRider, nDataRider)
				fmt.Print("Pilih data yang ingin dihapus dari NO rider: ")
				fmt.Scan(&x)
				idx = cariNoIdx(dataRider, nDataRider, x)
				if idx == -1 {
					clear_screen()
					fmt.Println("Data tidak ditemukan")
				} else {
					fmt.Print("Anda yakin ingin hapus data tersebut (y/n)? ")
					fmt.Scan(&pilih)
					if pilih == "y" || pilih ==  "Y" {
						hapusData(&dataRider, &nDataRider, x)
						clear_screen()
						fmt.Println("Data berhasil di hapus.")
					} 
				}
				if pilih == "n" || pilih == "N" || pilih == "y" || pilih ==  "Y"  {
					break
				}
			}
			
		}
		fmt.Print("Anda masih ingin hapus (y/n)? ")
		fmt.Scan(&pilih)
		if pilih == "y" || pilih ==  "Y" {
			clear_screen()
		} 
		clear_screen()
		if pilih == "n" || pilih == "N" {
			break
		}
	}
}

func subMenuHapusSemuaRider() {
	var bin, pilih string

	clear_screen()
	if nDataRider == 0 {
		fmt.Println("Data rider masih kosong.")
		fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
	} else {
		for {
			clear_screen()
			fmt.Print("Anda yakin ingin hapus semua data tersebut (y/n)? ")
			fmt.Scan(&pilih)
			if pilih == "y" || pilih ==  "Y" {
				resetData(&dataRider, &nDataRider)
				clear_screen()
				fmt.Println("Data berhasil di hapus.")
			}
			if pilih == "n" || pilih == "N" || pilih == "y" || pilih ==  "Y"  {
				break
			}
		}
	}
	fmt.Print("Tekan enter untuk keluar. ")
	fmt.Scanln(&bin)
}

/////////////////////////// Sub Menu Tampilan Edit  ///////////////////////////
func subMenuEditRider() {
	var pilih string
	var x, idx int

	for {
		clear_screen()
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			for {
				cetakRider(dataRider, nDataRider)
				fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
				fmt.Scan(&x)
				idx = cariNoIdx(dataRider, nDataRider, x)
				if idx == -1 {
					clear_screen()
					fmt.Println("Data tidak ditemukan")
				} else {
					fmt.Print("Anda yakin ingin hapus data tersebut (y/n)? ")
					fmt.Scan(&pilih)
					if pilih == "y" || pilih ==  "Y" {
						// lakukan edit data
						fmt.Print("Masukkan  nama ")
						clear_screen()
						fmt.Println("Data berhasil di hapus.")
					} 
				}
				if pilih == "n" || pilih == "N" || pilih == "y" || pilih ==  "Y"  {
					break
				}
			}
			
		}
		fmt.Print("Anda masih ingin hapus (y/n)? ")
		fmt.Scan(&pilih)
		if pilih == "y" || pilih ==  "Y" {
			clear_screen()
		} 
		clear_screen()
		if pilih == "n" || pilih == "N" {
			break
		}
	}
}

/////////////////////////// Input Data ///////////////////////////
func inputRider(T *tabRider, n *int, p string) {
	var i, min, sec, milsec, banyak, nAwal int
	var bin string

	clear_screen()
	if p == "1"{
		// mengecek apakah data sudah penuh 
		if *n == 24 {
			fmt.Println("Data sudah penuh.")
			fmt.Print("Tekan enter untuk kembali. ")
			fmt.Scanln(&bin)
		} else {
			// Proses vadidasi banyaknya rider
			nAwal = *n
			for {
				fmt.Print("Masukkan jumlah rider(1-", NMAX - *n, "): ")
				fmt.Scan(&banyak)
				if !(1 <= banyak && banyak <= NMAX - nAwal) {
					clear_screen()
					fmt.Println("Masukan tidak valid")
				} else {
					*n += banyak
				}
				if 1 <= banyak && banyak <= NMAX - nAwal {
					break
				}
			}
			// proses input data
			for i = *n - banyak; i < *n; i++ {
				fmt.Print("Input Data ridder ke-", i+1, "\n")
				fmt.Print("Input nomor rider: ")
				fmt.Scan(&T[i].no)
				fmt.Print("Input nama rider: ")
				inputFrasa(&T[i].name)
				fmt.Print("Input asal negara rider: ")
				fmt.Scan(&T[i].nat)
				fmt.Print("Input tim rider: ")
				inputFrasa(&T[i].team)
				fmt.Println()
			}
		}
	} else if p == "2" {
		if cekKelengkapanRider(*T, *n){
			urutNoRider(T, *n)
			fmt.Println("Input waktu FP1 rider dengan format [menit detik milidetik]:")
			for i = 0; i < *n; i++ {
				fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
				fmt.Scan(&min, &sec, &milsec)
				T[i].FP1 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "3" {
		// mengecek apakah data sudah penuh 
		if cekKelengkapanRider(*T , *n){
			fmt.Println("Data sudah penuh.")
			fmt.Print("Tekan enter untuk kembali. ")
			fmt.Scanln(&bin)
		}
		if cekKelengkapanRider(*T , *n) {
			urutNoRider(T, *n)
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu PR rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].PR = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "4" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu FP2 rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].FP2 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "5" {
		if cekKelengkapanRider(*T , *n) {
			urutWaktuPR(T, *n)
			for i = 12; i < *n; i++ {
				fmt.Println("Input waktu Q1 rider",  T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].Q1 = (60000 * min) + (1000 * sec) + milsec
			}
		}	
	} else if p == "6" {
		if cekKelengkapanRider(*T , *n) {
			urutWaktuPR(T, *n)
			for i = 0; i < 10; i++ {
				fmt.Println("Input waktu Q2 rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].Q2 = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "7" {
		if cekKelengkapanRider(*T , *n) {
			urutNoRider(T, *n)
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu SPR rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].SPR = (60000 * min) + (1000 * sec) + milsec
			}
		}
	} else if p == "8" {
		if cekKelengkapanRider(*T , *n) {
			urutNoRider(T, *n)
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu WUP rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].WUP = (60000 * min) + (1000 * sec) + milsec
			}
		}
		
	} else if p == "9" {
		if cekKelengkapanRider(*T , *n) {
			for i = 0; i < *n; i++ {
				fmt.Println("Input waktu RAC rider", T[i].name, "dengan format [menit detik milidetik]:")
				fmt.Scan(&min, &sec, &milsec)
				T[i].RAC = (60000 * min) + (1000 * sec) + milsec 
			}
		}
	}
}

/////////////////////////// Edit Data ///////////////////////////

/////////////////////////// Fungsi ///////////////////////////
func konversiWaktuTampilan(t, t0 , x int) string {
/*mengembalikan string dengan format “(menit)’(detik).(miliDetik)s” 
Jika posisi 1. Jika posisi > 1, maka mengembalikan string jarak dengan posisi pertama dengan. 
Jika waktu negatif, maka mengembalikan perolehan laps terakhir*/
	var menit, detik, miliDetik int
	var waktu string

	if t == -1 {
		waktu = "DNF"
	} else if x == 0  {
		menit = t / 60000
		t %= 60000
		detik = t / 1000
		t %= 1000
		miliDetik = t
		waktu = strconv.Itoa(menit) + "'" + strconv.Itoa(detik) + "." + strconv.Itoa(miliDetik) + "s"
	} else {
		t -= t0
		detik = t / 1000
		t %= 1000
		miliDetik = t
		waktu = "+" + strconv.Itoa(detik) + "." + strconv.Itoa(miliDetik) + "s"
	}
	return waktu
}

func hitungTotalPoin(T *tabRider, n int) {
/*I.S.
F.S.  */
	var i int

    urutWaktuSPR(T, n)
	T[0].poinSPR = 12
	T[1].poinSPR = 9
	T[2].poinSPR = 7
	T[3].poinSPR = 6
	T[4].poinSPR = 5
	T[5].poinSPR = 4
	T[6].poinSPR = 3
	T[7].poinSPR = 2
	T[8].poinSPR = 1
	for i = n; i < NMAX; i++ {
		T[i].poinSPR = 0
	}
    urutWaktuRAC(T, n)
	T[0].poinRAC = 25
	T[1].poinRAC = 20
	T[2].poinRAC = 16
	T[3].poinRAC = 13
	T[4].poinRAC = 11
	T[5].poinRAC = 10
	T[6].poinRAC = 9
	T[7].poinRAC = 8
	T[8].poinRAC = 7
	T[9].poinRAC = 6
	T[10].poinRAC = 5
	T[11].poinRAC = 4
	T[12].poinRAC = 3
	T[13].poinRAC = 2
	T[14].poinRAC = 1
	for i = n; i < NMAX; i++ {
		T[i].poinRAC = 0
	}
	for i = 0; i < n; i++ {
		T[i].totalPoin = T[i].poinSPR + T[i].poinRAC
	}
}

func cariNoIdx(T tabRider, n, x int) int {
	var i, ketemu int

	ketemu = -1
	i = 0
	for i < n && ketemu == -1 {
		if x == T[i].no {
			ketemu = i
		}
		i++ 
	}
	return ketemu
}
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
		for i > 0 && temp.totalPoin > T[i - 1].totalPoin {
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

// func urutRaceGrid(T *tabRider, n int) {
// 	/*I.S. terdefinisi array acak T1 sebanyak n
// 	F.S. T berisi nilai terurut berdasarkan waktu Q2*/
// 	var i int
// 	var S tabRider

// 	urutWaktuPR(T, n)
// 	if n > 10 {
// 		for i = 10; i < n; i++ {
// 			S[i].Q1 = -1
// 		}
// 	}
// 	urutWaktuQ1(T, n)

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

/////////////////////////// Cetak Tabel ///////////////////////////
func cetakRider(T tabRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team*/
	var i int

	urutNoRider(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v\n", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v\n", T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakRaceGrid(T tabRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team*/
	var i int

	// urut berdasarkan Q1 dan Q2

	// lakukan validasi bahwa data Q1 dan Q2 lengkap

	fmt.Printf("%-25v %-2v %3v %25v\n", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25s %-2v %3s %25s\n", T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakWaktuFP1(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuFP1(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP1, T[0].FP1, i))
	}
}

func cetakWaktuPR(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuPR(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].PR, T[0].PR, i))
	}
}

func cetakWaktuFP2(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuFP2(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP2, T[0].FP2, i))
	}
}

func cetakWaktuQ1(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuQ1(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q1, T[0].Q1, i))
	}
}

func cetakWaktuQ2(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuQ2(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q2, T[0].Q2, i))
	}
}

func cetakWaktuSPR(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuSPR(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].SPR, T[0].SPR, i))
	}
}

func cetakWaktuWUP(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuWUP(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].WUP, T[0].WUP, i))
	}
}

func cetakWaktuRAC(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutWaktuRAC(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].RAC, T[0].RAC, i))
	}
}

func cetakPoin(T tabRider, n int) {
/*I.S T array dan banyak elemen n terdefinisi
F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
	var i int

	urutTotalPoin(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-8v %-8v %-4v\n", "NAME", "NO", "NAT", "TEAM", "POIN SPR", "POIN RAC", "POIN")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-8v %-8v %-4v\n", T[i].name, T[i].no, T[i].nat, T[i].team, T[i].poinSPR, T[i].poinRAC, T[i].totalPoin)
	}
}

/////////////////////////// Edit ///////////////////////////


/////////////////////////// Hapus ///////////////////////////
func hapusData(T *tabRider, n *int, x int) {
	var i int

	for i = x; i < *n - 1; i++ {
		T[i] = T[i + 1]
	} 
	hapusBagian(T, i)
	*n--
	hitungTotalPoin(T,*n)
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

/////////////////////////// Test data ///////////////////////////
func test(T *tabRider, n *int) {
	*T = tabRider{
		{no: 1, name: "Jorge Martin", nat: "USA", team: "TeamA", FP1: 121, PR: 241, FP2: 314, Q1: 415, Q2: 265, SPR: 662, WUP: 147, RAC: 238, poinSPR: 160, poinRAC: 260, totalPoin: 320},
		{no: 2, name: "Marc Marquez", nat: "ESP", team: "TeamB", FP1: 223, PR: 324, FP2: 451, Q1: 552, Q2: 266, SPR: 257, WUP: 158, RAC: 619, poinSPR: 215, poinRAC: 325, totalPoin: 820},
		{no: 3, name: "Maverick Viñales", nat: "ESP", team: "TeamB", FP1: 242, PR: 311, FP2: 834, Q1: 535, Q2: 236, SPR: 737, WUP: 871, RAC: 249, poinSPR: 155, poinRAC: 265, totalPoin: 371},
		{no: 4, name: "Enea Bastianini", nat: "ESP", team: "TeamB", FP1: 215, PR: 314, FP2: 415, Q1: 527, Q2: 226, SPR: 617, WUP: 158, RAC: 679, poinSPR: 115, poinRAC: 235, totalPoin: 621},
		{no: 5, name: "Aleix Espargaro", nat: "ESP", team: "TeamB", FP1: 142, PR: 341, FP2: 444, Q1: 512, Q2: 726, SPR: 176, WUP: 826, RAC: 219, poinSPR: 125, poinRAC: 265, totalPoin: 521},
		{no: 6, name: "Pedro Acosta", nat: "ESP", team: "TeamB", FP1: 214, PR: 123, FP2: 414, Q1: 514, Q2: 286, SPR: 721, WUP: 826, RAC: 629, poinSPR: 165, poinRAC: 235, totalPoin: 141},
	}
	*n = 6
}
/*
no                                  int
name, nat, team                     string
FP1, PR, FP2, Q1, Q2, SPR, WUP, RAC int
poinSPR, poinRAC, totalPoin
*/ 