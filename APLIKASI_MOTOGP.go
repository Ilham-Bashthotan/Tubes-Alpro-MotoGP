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
	for {
		clear_screen()
		intro()
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
		inputFrasa(&pilih)
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

/////////////////////////// Tampilan intro dan outro ///////////////////////////
func intro() {
	/*
		IS: -
		FS: Tercetak tampilan intro aplikasi MotoGP
	*/
    clear_screen()
    fmt.Println("Selamat datang di aplikasi")
	fmt.Println("______  ___     _____      ________________")
	fmt.Println("___   |/  /_______  /________  ____/__  __ \\")
	fmt.Println("__  /|_/ /_  __ \\  __/  __ \\  / __ __  /_/ /")
	fmt.Println("_  /  / / / /_/ / /_ / /_/ / /_/ / _  ____/")
	fmt.Println("/_/  /_/  \\____/\\__/ \\____/\\____/  /_/")
}

func bye() {
	/*
		IS: -
		FS: Tercetak tampilan outro aplikasi MotoGP
	*/
    clear_screen()
    fmt.Println("Terima kasih telah menggunakan aplikasi MotoGP")
    fmt.Println("By:")
    fmt.Println("Ilham Bashthotan")
    fmt.Println("Faiz Nawfal Zulfiqar")
}

///////////////////////////  tampilan menu ///////////////////////////
func inputDataRider() {
	/*
		IS: -
		FS: Tercetak tampilan menu input data rider
	*/
	var pilih string

	for {
		clear_screen()
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
		inputFrasa(&pilih)
		switch pilih {
			case "1": subMenuInput(&dataRider, &nDataRider, "Rider")
			case "2": subMenuInput(&dataRider, &nDataRider, "FP1")
			case "3": subMenuInput(&dataRider, &nDataRider, "PR")
			case "4": subMenuInput(&dataRider, &nDataRider, "FP2")
			case "5": subMenuInput(&dataRider, &nDataRider, "Q1")
			case "6": subMenuInput(&dataRider, &nDataRider, "Q2")
			case "7": subMenuInput(&dataRider, &nDataRider, "SPR")
			case "8": subMenuInput(&dataRider, &nDataRider, "WUP")
			case "9": subMenuInput(&dataRider, &nDataRider, "RAC")
			default: clear_screen()
		}
		if pilih == "10" {
			break
		}
	}
}

func tampilkanDataRider() {
	/*
		IS: -
		FS: Tercetak tampilan menu tampilkan data rider
	*/
	var pilih string

	for {
		clear_screen()
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
		fmt.Println("11. Total Poin")
		fmt.Println("12. Kembali")
		fmt.Println("--------------------------")
		fmt.Print("Pilih (1-11): ")
		inputFrasa(&pilih)
		switch pilih {
			case "1":  subMenuCetak(dataRider, nDataRider, "Rider")
			case "2":  subMenuCetak(dataRider, nDataRider, "RaceGrid")
			case "3":  subMenuCetak(dataRider, nDataRider, "FP1")
			case "4":  subMenuCetak(dataRider, nDataRider, "PR")
			case "5":  subMenuCetak(dataRider, nDataRider, "FP2")
			case "6":  subMenuCetak(dataRider, nDataRider, "Q1")
			case "7":  subMenuCetak(dataRider, nDataRider, "Q2")
			case "8":  subMenuCetak(dataRider, nDataRider, "SPR")
			case "9":  subMenuCetak(dataRider, nDataRider, "WUP")
			case "10": subMenuCetak(dataRider, nDataRider, "RAC")
			case "11": subMenuCetak(dataRider, nDataRider, "Poin")
			default: clear_screen()
		}
		if pilih == "12" {
			break
		}
	}
}

func editDataRider() {
	/*
		IS: -
		FS: Tercetak tampilan menu edit data rider
	*/
	var pilih string

	for {
		clear_screen()
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
		inputFrasa(&pilih)
		switch pilih {
			case "1": subMenuEdit(&dataRider, nDataRider, "Rider")
			case "2": subMenuEdit(&dataRider, nDataRider, "FP1")
			case "3": subMenuEdit(&dataRider, nDataRider, "PR")
			case "4": subMenuEdit(&dataRider, nDataRider, "FP2")
			case "5": subMenuEdit(&dataRider, nDataRider, "Q1")
			case "6": subMenuEdit(&dataRider, nDataRider, "Q2")
			case "7": subMenuEdit(&dataRider, nDataRider, "SPR")
			case "8": subMenuEdit(&dataRider, nDataRider, "WUP")
			case "9": subMenuEdit(&dataRider, nDataRider, "RAC")
			default: clear_screen()
		}
		if pilih == "10" {
			break
		}
	}
}

func hapusDataRider() {
	/*
		IS: -
		FS: Tercetak tampilan menu hapus data rider
	*/
	var pilih string

	for {
		clear_screen()
		fmt.Println("--------------------------")
		fmt.Println("     Hapus Data Rider     ")
		fmt.Println("--------------------------")
		fmt.Println("1. Hapus Rider")
		fmt.Println("2. Hapus Semua Rider")
		fmt.Println("3. Kembali")
		fmt.Println("--------------------------")
		fmt.Print("Pilih (1/2/3): ")
		inputFrasa(&pilih)
		switch pilih {
			case "1": subMenuHapus(&dataRider, &nDataRider, "Rider")
			case "2": subMenuHapus(&dataRider, &nDataRider, "All")
			default: clear_screen()
		}
		if pilih == "3" {
			break
		}
	}
}

/////////////////////////// Sub Menu ///////////////////////////
func subMenuInput(T *tabRider, n *int, p string) {
	/*
		IS: Terdefinisi T array sebanyak n, terdefinisi p pilihan yang merujuk pada field pada T array
		FS: array T terisi sesuai p pilihan pada field array T, n tersi banyaknya T rider yang sesuai pada p pilihan
	*/
	var i, min, sec, milsec, banyak, nAwal, no int
	var bin string

	clear_screen()
	if p == "Rider"{
		// mengecek apakah data sudah penuh 
		if *n == 24 {
			fmt.Println("Data sudah penuh.")
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
				fmt.Scan(&no)
				for member(*T, *n, no) {
					fmt.Println("Nomor rider dengan no", no, "sudah ada.")
					fmt.Println("Coba inputkan nomor rider yang lain")
					fmt.Print("Input nomor rider: ")
					fmt.Scan(&no)
				}
				T[i].no = no
				fmt.Print("Input nama rider: ")
				inputFrasa(&T[i].name)
				fmt.Print("Input asal negara rider: ")
				fmt.Scan(&T[i].nat)
				fmt.Print("Input tim rider: ")
				inputFrasa(&T[i].team)
				fmt.Println()
			}
			clear_screen()
			fmt.Println("Data Rider berhasil diinput. ")
		}
	} else if p == "FP1" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanFP1(*T, *n) {
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu FP1 rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].FP1 == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].FP1 = konversiWaktuMs(min, sec, milsec)
				}
			}
			clear_screen()
			fmt.Println("Data FP1 berhasil diinput. ")
		}
	} else if p == "PR" {
		// mengecek apakah data sudah penuh 
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanPR(*T , *n){
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu PR rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].PR == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].PR = konversiWaktuMs(min, sec, milsec)
				}
			}
			// pengisian untuk rider yang tidak masuk ke Q1
			urutWaktuPR(T, *n)
			i = 0
			for i < 10 && i < *n  {
				T[i].Q1 = -1
				i++
			}
			// pengisian untuk rider yang tidak masuk ke Q2
			for i = 10; i < *n; i++ {
				if T[i].Q2 <= 0 {
					T[i].Q2 = -1
				}
			}
			clear_screen()
			fmt.Println("Data PR berhasil diinput. ")
		}
	} else if p == "FP2" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanFP2(*T, *n) {
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu FP2 rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].FP2 == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].FP2 = konversiWaktuMs(min, sec, milsec)
				}
			}
			clear_screen()
			fmt.Println("Data FP2 berhasil diinput. ")
		}
	} else if p == "Q1" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanQ1(*T, *n) {
			fmt.Println("Data sudah penuh.")
		} else if !cekKelengkapanPR(*T, *n) {
			fmt.Println("Data PR masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu Q1 rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].Q1 == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].Q1 = konversiWaktuMs(min, sec, milsec)
				}
			}
			// pengisian untuk rider yang masuk ke Q2
			urutWaktuQ1(T, *n)
			i = 0
			for i < 2 && i < *n  {
				if T[i].Q2 < 0 {
					T[i].Q2 = 0
				}
				i++
			}
			clear_screen()
			fmt.Println("Data Q1 berhasil diinput. ")
		}
	} else if p == "Q2" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanQ2(*T , *n) {
			fmt.Println("Data sudah penuh.")
		} else if !cekKelengkapanQ1(*T, *n) {
			fmt.Println("Data Q1 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu Q2 rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].Q2 == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].Q2 = konversiWaktuMs(min, sec, milsec)
				}
			}
			clear_screen()
			fmt.Println("Data Q2 berhasil diinput. ")
		}
	} else if p == "SPR" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanSPR(*T , *n) {
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu SPR rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].SPR == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].SPR = konversiWaktuMs(min, sec, milsec)
				}
			}
			clear_screen()
			fmt.Println("Data SPR berhasil diinput. ")
		}
	} else if p == "WUP" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanWUP(*T , *n) {
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu WUP rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].WUP == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].WUP = konversiWaktuMs(min, sec, milsec)
				}
			}
			clear_screen()
			fmt.Println("Data WUP berhasil diinput. ")
		}
		
	} else if p == "RAC" {
		if nDataRider == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if cekKelengkapanRAC(*T , *n) {
			fmt.Println("Data sudah penuh.")
		} else {
			urutNoRider(T, *n)
			fmt.Println("Input waktu RAC rider dengan format [menit detik milidetik]: ")
			for i = 0; i < *n; i++ {
				if T[i].RAC == 0 {
					fmt.Print("#", T[i].no, " ",  T[i].name, ": ")
					fmt.Scan(&min, &sec, &milsec)
					T[i].RAC = konversiWaktuMs(min, sec, milsec) 
				}
			}
			clear_screen()
			fmt.Println("Data RAC berhasil diinput. ")
		}
	}
	fmt.Print("Tekan enter untuk kembali. ")
	inputFrasa(&bin)
}

func subMenuCetak(T tabRider, n int, p string) {
	/*
		IS: Terdefinisi T array sebanyak n, terdefinisi p pilihan yang merujuk pada field pada T array atau pilihan di menu cetak
		FS: cetak data yang tertapat pada array T yang sesuai pada p pilihan
	*/
	var bin string

	clear_screen()
	if p == "Rider" {	
		fmt.Println("--------------------------")
		fmt.Println("        Data Rider        ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakRider(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)
	} else if p == "RaceGrid" {
		fmt.Println("--------------------------")
		fmt.Println("      Data Race Grid      ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data Rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !(cekKelengkapanQ1(T, n) && cekKelengkapanQ2(T, n)) {
			fmt.Println("Data Q1 atau Q2 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")		
		} else {
			cetakRaceGrid(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	} else if p == "Poin" {
		fmt.Println("--------------------------")
		fmt.Println("        Data Poin         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data Rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !(cekKelengkapanQ1(T, n) && cekKelengkapanQ2(T, n)) {
			fmt.Println("Data Q1 atau Q2 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			hitungTotalPoin(&T, n)
			cetakPoin(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)
	} else if p == "FP1" {
		fmt.Println("--------------------------")
		fmt.Println("         Data FP1         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanFP1(T, n) {
			fmt.Println("Data FP1 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuFP1(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)
		fmt.Print("Tekan enter sekali lagi. ")
	} else if p == "PR" {
		fmt.Println("--------------------------")
		fmt.Println("         Data PR          ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanPR(T, n) {
			fmt.Println("Data PR masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuPR(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)
	} else if p == "FP2" {
		fmt.Println("--------------------------")
		fmt.Println("         Data FP2         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanFP2(T, n) {
			fmt.Println("Data FP2 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuFP2(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	} else if p == "Q1" {
		fmt.Println("--------------------------")
		fmt.Println("         Data Q1          ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanQ1(T, n) {
			fmt.Println("Data Q1 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuQ1(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	} else if p == "Q2" {
		var bin string

		fmt.Println("--------------------------")
		fmt.Println("         Data Q2          ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanQ2(T, n) {
			fmt.Println("Data Q2 masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuQ2(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	} else if p == "SPR" {
		fmt.Println("--------------------------")
		fmt.Println("         Data SPR         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanSPR(T, n) {
			fmt.Println("Data SPR masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuSPR(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	} else if p == "WUP" {
		fmt.Println("--------------------------")
		fmt.Println("         Data WUP         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanWUP(T, n) {
			fmt.Println("Data WUP masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuWUP(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)
	} else if p == "RAC" {
		fmt.Println("--------------------------")
		fmt.Println("         Data RAC         ")
		fmt.Println("--------------------------")
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else if !cekKelengkapanRAC(T, n) {
			fmt.Println("Data RAC masih belum lengkap.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			cetakWaktuRAC(T, n)
		}
		fmt.Print("Tekan enter untuk keluar. ")
		fmt.Scanln(&bin)		
	}
}

func subMenuHapus(T *tabRider, n *int, p string) {
	/*
		IS: Terdefinisi T array sebanyak n, terdefinisi p pilihan yang merujuk pada field pada T array
		FS: hapus data yang tertapat pada array T yang sesuai pada p pilihan
	*/	
	var bin, pilih string
	var x, idx int

	if p == "Rider" {
		for {
			clear_screen()
			if nDataRider == 0 {
				fmt.Println("Data rider masih kosong.")
				fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
			} else {
				for {
					cetakRider(*T, *n)
					fmt.Print("Pilih data yang ingin dihapus dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, *n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan")
					} else {
						fmt.Print("Anda yakin ingin hapus data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							hapusData(T, n, idx)
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
	} else if p == "All" {
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
}

func subMenuEdit(T *tabRider, n int, p string) {
	/*
		IS: Terdefinisi T array sebanyak n, terdefinisi p pilihan yang merujuk pada field pada T array
		FS: hapus data yang tertapat pada array T yang sesuai pada p pilihan
	*/
	var pilih string
	var x, idx, i int
	var no int
	var name, nat, team string
	var min, sec, milsec int

	for {
		clear_screen()
		if n == 0 {
			fmt.Println("Data rider masih kosong.")
			fmt.Println("Silakan isi terlebih dahulu di menu Input Data.")
		} else {
			for {
				if p == "Rider" {
					// cetak tabel
					cetakRider(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Input nomor rider: ")
						fmt.Scan(&no)
						for member(*T, n, no) {
							fmt.Println("Nomor rider dengan no", no, "sudah ada.")
							fmt.Println("Coba inputkan nomor rider yang lain")
							fmt.Print("Input nomor rider: ")
							fmt.Scan(&no)
						}
						fmt.Print("Input nama rider: ")
						inputFrasa(&name)
						fmt.Print("Input asal negara rider: ")
						fmt.Scan(&nat)
						fmt.Print("Input tim rider: ")
						inputFrasa(&team)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							T[idx].no = no
							T[idx].name = name
							T[idx].nat = nat
							T[idx].team = team
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "FP1" {
					// cetak tabel
					cetakWaktuFP1(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Edit waktu FP1 rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].FP1 = konversiWaktuMs(min, sec, milsec)
							// pengisian untuk rider yang tidak masuk ke Q1
							urutWaktuPR(T, n)
							i = 0
							for i < 10 && i < n  {
								T[i].Q1 = -1
								i++
							}
							// pengisian untuk rider yang tidak masuk ke Q2
							for i = 10; i < n; i++ {
								if T[i].Q2 <= 0 {
									T[i].Q2 = -1
								}
							}
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "PR" {
					// cetak tabel
					cetakWaktuPR(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Edit waktu PR rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].PR = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
							// pengisian untuk yang tidak masuk ke Q2
							urutWaktuQ1(T, n)
							i = 0
							for 2 <= i && i <= 11 && i < n  {
								T[i].Q2 = -1
								i++
							}
							// pengisian untuk rider yang masuk dari Q1 ke Q2
							i = 0
							for i < 2 && i < n  {
								if T[i].Q2 <= 0 {
									T[i].Q2 = 0
									i++
								}
							}
						} 
					}
				} else if p == "FP2" {
					// cetak tabel
					cetakWaktuFP2(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Edit waktu FP2 rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].FP2 = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "Q1" {
					// cetak tabel
					cetakWaktuQ1(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 || T[idx].Q1 <= 0 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Edit waktu Q1 rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].Q1 = konversiWaktuMs(min, sec, milsec)
							// pengisian untuk rider yang masuk ke Q2
							urutWaktuQ1(T, n)
							i = 0
							for i < 2 && i < n  {
								if T[i].Q2 < 0 {
									T[i].Q2 = 0
								}
								i++
							}
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "Q2" {
					// cetak tabel
					cetakWaktuQ2(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 || T[idx].Q1 <= 0 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {
						// lakukan edit data
						fmt.Print("Edit waktu Q2 rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].Q2 = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "SPR" {
					// cetak tabel
					cetakWaktuSPR(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {						
						// lakukan edit data
						fmt.Print("Edit waktu SPR rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].SPR = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "WUP" {
					// cetak tabel
					cetakWaktuWUP(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {						
						// lakukan edit data
						fmt.Print("Edit waktu WUP rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].WUP = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} else if p == "RAC" {
					// cetak tabel
					cetakWaktuRAC(*T, n)
					fmt.Print("Pilih rider yang ingin diedit dari NO rider: ")
					fmt.Scan(&x)
					idx = cariNoIdx(*T, n, x)
					if idx == -1 {
						clear_screen()
						fmt.Println("Data tidak ditemukan.")
					} else {						
						// lakukan edit data
						fmt.Print("Edit waktu RAC rider dengan format [menit detik milidetik]: ")
						fmt.Scan(&min, &sec, &milsec)
						fmt.Println()
						// tanyakan kembali
						fmt.Print("Anda yakin ingin mengedit data tersebut (y/n)? ")
						fmt.Scan(&pilih)
						if pilih == "y" || pilih ==  "Y" {
							clear_screen()
							dataRider[idx].RAC = konversiWaktuMs(min, sec, milsec)
							fmt.Println("Data berhasil diedit.")
						} 
					}
				} 
				if pilih == "n" || pilih == "N" || pilih == "y" || pilih ==  "Y"  {
					break
				}
			}
			
		}
		fmt.Print("Anda masih ingin edit (y/n)? ")
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

/////////////////////////// Hitung Poin ///////////////////////////
func hitungTotalPoin(T *tabRider, n int) {
	/*
		IS: terdefinisi T array sebanyak n
		FS: total poin terisi pada array T
	*/
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

/////////////////////////// Urut Data ///////////////////////////
func urutNoRider(T *tabRider, n int) {
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu no
	*/
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
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu FP1
	*/
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
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu PR
	*/
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
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu FP2
	*/
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
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu Q1
	*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && ((temp.Q1 < T[i - 1].Q1 || T[i - 1].Q1 == -1) && temp.Q1 != -1) {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuQ2(T *tabRider, n int) {
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu Q2
	*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && ((temp.Q2 < T[i - 1].Q2 || T[i - 1].Q2 == -1) && temp.Q2 != -1) {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuSPR(T *tabRider, n int) {
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu SPR
	*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && ((temp.SPR < T[i - 1].SPR || T[i - 1].SPR == -1) && temp.SPR != -1) {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutWaktuWUP(T *tabRider, n int) {
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu WUP
	*/
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
/*
	IS: terdefinisi array T sebanyak n
	FS: T berisi nilai terurut berdasarkan waktu totalPoin*/
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
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu RAC
	*/
	var pass, i int
	var temp rider

	for pass = 1; pass < n; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && ((temp.RAC < T[i - 1].RAC || T[i - 1].RAC == -1) && temp.RAC != -1) {
			T[i] = T[i - 1]
			i--
		}
		T[i] = temp
	}
}

func urutRaceGrid(T *tabRider, n int) {
	/*
		IS: terdefinisi array T sebanyak n
		FS: T berisi nilai terurut berdasarkan waktu Q2
	*/
	urutWaktuQ1(T, n)
	urutWaktuQ2(T, n)
}

/////////////////////////// Cetak Tabel ///////////////////////////
func cetakRider(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team
	*/
	var i int

	urutNoRider(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v\n", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v\n", T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakRaceGrid(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team
	*/
	var i int

	urutRaceGrid(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v\n", "POS", "NAME", "NO", "NAT", "TEAM")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-25s %-2v %-3v %-25v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team)
	}
}

func cetakWaktuFP1(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, FP1
	*/
	var i int

	urutWaktuFP1(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP1, T[0].FP1, i))
	}
}

func cetakWaktuPR(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, PR
	*/
	var i int

	urutWaktuPR(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].PR, T[0].PR, i))
	}
}

func cetakWaktuFP2(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, FP2
	*/
	var i int

	urutWaktuFP2(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].FP2, T[0].FP2, i))
	}
}

func cetakWaktuQ1(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, Q1
	*/
	var i int

	urutWaktuQ1(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		if T[i].Q1 > 0 {
			fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n",  i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q1, T[0].Q1, i))
		}
	}
}

func cetakWaktuQ2(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, Q2
	*/
	var i int

	urutWaktuQ2(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		if T[i].Q2 > 0 {
			fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].Q2, T[0].Q2, i))
		}
	}
}

func cetakWaktuSPR(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, SPR
	*/
	var i int

	urutWaktuSPR(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		if T[i].SPR > 0 {
			fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].SPR, T[0].SPR, i))
		} else {
			fmt.Printf("    %-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].SPR, T[0].SPR, i))
		}
	}
}

func cetakWaktuWUP(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, WUP
	*/
	var i int

	urutWaktuWUP(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].WUP, T[0].WUP, i))
	}
}

func cetakWaktuRAC(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, RAC
	*/
	var i int

	urutWaktuRAC(&T, n)
	fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", "POS", "NAME", "NO", "NAT", "TEAM", "TIME/DIFF")
	for i = 0; i < n; i++ {
		if T[i].RAC > 0 {
			fmt.Printf("%-3v %-25v %-2v %-3v %-25v %-11v\n", i + 1, T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].RAC, T[0].RAC, i))
		} else {
			fmt.Printf("    %-25v %-2v %-3v %-25v %-11v\n", T[i].name, T[i].no, T[i].nat, T[i].team, konversiWaktuTampilan(T[i].RAC, T[0].RAC, i))
		}
	}
}

func cetakPoin(T tabRider, n int) {
	/*
		IS: T array dan banyak elemen n terdefinisi
		FS: tercetak tabel dari field : name, no, nat, team, totalPoin
	*/
	var i int

	urutTotalPoin(&T, n)
	fmt.Printf("%-25v %-2v %-3v %-25v %-8v %-8v %-4v\n", "NAME", "NO", "NAT", "TEAM", "POIN SPR", "POIN RAC", "POIN")
	for i = 0; i < n; i++ {
		fmt.Printf("%-25v %-2v %-3v %-25v %-8v %-8v %-4v\n", T[i].name, T[i].no, T[i].nat, T[i].team, T[i].poinSPR, T[i].poinRAC, T[i].totalPoin)
	}
}

/////////////////////////// Hapus data ///////////////////////////
func hapusData(T *tabRider, n *int, x int) {
	/*
		IS: Array T dan jumlah data n terdefinisi sembarang, index x terdefinisi
		FS: Data array T yang ada pada index x terhapus dan data semua pada index setelahnya bergeser ke kiri
	*/
	var i int

	for i = x; i < *n - 1; i++ {
		T[i] = T[i + 1]
	} 
	hapusBagian(T, i)
	*n--
	// Proses menghitung poin
	hitungTotalPoin(T,*n)
	// pengisian untuk rider yang tidak masuk ke Q1
	urutWaktuPR(T, *n)
	i = 0
	for i < 10 && i < *n  {
		T[i].Q1 = -1
		i++
	}
	// pengisian untuk rider yang tidak masuk ke Q2
	for i = 10; i < *n; i++ {
		if T[i].Q2 <= 0 {
			T[i].Q2 = -1
		}
	}
	// pengisian untuk rider yang masuk ke Q2
	urutWaktuQ1(T, *n)
	i = 0
	for i < 2 && i < *n  {
		if T[i].Q2 < 0 {
			T[i].Q2 = 0
		}
		i++
	}
}

func resetData(T *tabRider, n *int) {
	/*
		IS: Array T dan jumlah data n terdefinisi sembarang
		FS: Semua data pad array T terhapus dan nilai n menjadi 0
	*/
	var i int

	for i = 0; i < NMAX; i++ {
		hapusBagian(T, i)
	} 
	*n = 0
}

func hapusBagian(T *tabRider, idx int) {
	/*
		IS: Array T terdefinisi sembarang, index idx terdefinisi
		FS: Semua data pad array T terhapus dan nilai n menjadi 0
	*/
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

/////////////////////////// function ///////////////////////////
func cekKelengkapanRider(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua nomor rider pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat nomor yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP1(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu FP1 pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu FP1 yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanPR(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu PR pada arrray 
		bentukan T sebanyan n data sudah terisi dan nilai false apablia 
		terdapat waktu PR yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].PR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP2(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu FP2 pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu FP2 yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ1(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu Q1 pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu Q1 yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ2(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu Q2 pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu Q2 yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanSPR(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true apabila semua waktu SPR pada array 
		bentuka T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu SPR yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].SPR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanWUP(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true jika semua waktu WUP pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu WUP yang bernilai 0 
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].WUP == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanRAC(T tabRider, n int) bool {
	/*
		Mengembalikan nilai true jika semua waktu RAC pada array 
		bentukan T sebanyak n data sudah terisi dan nilai false apabila 
		terdapat waktu RAC yang bernilai 0
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].RAC == 0 {
			return false
		}
	} 
	return true
}

func member(T tabRider, n int, x int) bool {
	/*
		Mengembalikan nilai true apabila ditemukan rider dengan nomor x
		dan false apabila tidak ditemukan rider dengan nomor x
	*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == x {
			return true
		}
	} 
	return false
}

func konversiWaktuTampilan(t, t0 , x int) string {
	/*
		mengembalikan string dengan format “(menit)’(detik).(miliDetik)s” 
		Jika posisi 1. Jika posisi > 1, maka mengembalikan string jarak dengan posisi pertama dengan. 
		Jika waktu negatif, maka mengembalikan perolehan laps terakhir
	*/
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

func cariNoIdx(T tabRider, n, x int) int {
	/*
		Mengembalikan alamat index dari rider dengan nomor x
		dan -1 apabila tidak ditemukan rider dengan nomor x
	*/	
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

func konversiWaktuMs(m, s, ms int) int {
	/*
		mengembalikan nilai konversi waktu ke milidetik jika 
		m (menit), s (detik), ms (milidetik) valid, selain dari itu 
		maka mengembalikan nilai -1
	*/
	if m >= 0 && 0 <= s && s <= 59 && 0 <= ms && ms<= 999 {
		return m * 60000 + s * 1000 + ms
	}
	return -1
}

/////////////////////////// Hiasan ///////////////////////////
func inputFrasa(str *string) {
	/*
		IS: terdefinisi sembarang str (string)
		FS: str terisi frasa dan berhanti ketika masukan berupa enter
	*/
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

func clear_screen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}