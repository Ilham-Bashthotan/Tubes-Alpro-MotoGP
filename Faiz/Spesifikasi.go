func cekKelengkapanRider(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua nomor rider pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat nomor yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP1(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu FP1 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu FP1 yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanPR(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu PR pada arrray bentukan T sebanyan n data sudah terisi
dan nilai false apablia terdapat waktu PR yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].PR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP2(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu FP2	pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu FP2 yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ1(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu Q1 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu Q1 yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ2(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu Q2 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu Q2 yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanSPR(T tabRider, n int) bool {
/*Mengembalikan nilai true apabila semua waktu SPR pada array bentuka T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu SPR yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].SPR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanWUP(T tabRider, n int) bool {
/*Mengembalikan nilai true jika semua waktu WUP pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu WUP yang bernilai 0*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].WUP == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanRAC(T tabRider, n int) bool {
/*{Mengembalikan nilai true jika semua waktu RAC pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu RAC yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].RAC == 0 {
			return false
		}
	} 
	return true
}

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
	FS: Data yag ada pada index idx terhapus
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

func member(T tabRider, n int, x int) bool {
/*Mengembalikan nilai true apabila ditemukan rider dengan nomor x
dan false apabila tidak ditemukan rider dengan nomor x*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == x {
			return true
		}
	} 
	return false
}

func cariNoIdx(T tabRider, n, x int) int {
/*Mengembalikan alamat index dari rider dengan nomor x
dan -1 apabila tidak ditemukan rider dengan nomor x*/	
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