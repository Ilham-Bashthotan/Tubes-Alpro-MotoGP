package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX = 20

type produk struct {
	barang string
	harga  int
}

type tabProduk [NMAX]produk

func main() {
	var dataProduk tabProduk
	var nDataProduk int
	inputProduk(&dataProduk, &nDataProduk)
	cetakProduk(dataProduk, nDataProduk)
}

func inputProduk(T *tabProduk, n *int) {
	clear_screen()
	fmt.Print("masukkan banyak produk yang di input (1-20): ")
	// validasi banyaknya produk yang akan di-input
	for {
		fmt.Scan(n)
		if !(1 <= *n && *n <= NMAX) {
			clear_screen()
			fmt.Println("Masukan tidak valid")
			fmt.Print("masukkan banyak produk yang di input (1-20): ")
		}
		if 1 <= *n && *n <= NMAX {
			break
		}
	}
	// proses input data
	for i := 0; i < *n; i++ {
		fmt.Print("Data ke-", i + 1, "\n")
		fmt.Print("masukkan nama: ")
		inputNama(&T[i].barang)
		fmt.Print("masukkan harga: ")
		fmt.Scan(&T[i].harga)
	}
}

func cetakProduk(T tabProduk, n int) {
	clear_screen()
	fmt.Println("Isi Produk: ")
	// fmt.Println(T)
	for i := 0; i < n; i++ {
		fmt.Println(i+1, T[i].barang, T[i].harga)
	}
}

func inputNama(str *string) {
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