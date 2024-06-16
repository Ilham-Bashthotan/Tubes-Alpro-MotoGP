func cekKelengkapanRider(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua nomor rider pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat nomor yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP1(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu FP1 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu FP1 yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanPR(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu PR pada arrray bentukan T sebanyan n data sudah terisi
dan nilai false apablia terdapat waktu PR yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].PR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP2(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu FP2	pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu FP2 yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ1(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu Q1 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu Q1 yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanQ2(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu Q2 pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu Q2 yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].Q2 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanSPR(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua waktu SPR pada array bentuka T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu SPR yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].SPR == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanWUP(T tabRider, n int) bool {
/*{Mengembalikan nilai true jika semua waktu WUP pada array bentukan T sebanyak n data sudah terisi
dan nilai false apabila terdapat waktu WUP yang bernilai 0}*/
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