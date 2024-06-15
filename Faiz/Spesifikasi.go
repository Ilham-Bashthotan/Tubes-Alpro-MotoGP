func cekKelengkapanRider(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua array nomor rider sudah terisi
dan nilai false apabila terdapat array nomor yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].no == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanFP1(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua array waktu FP1 sudah terisi
dan nilai false apabila terdapat array waktu FP1 yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].FP1 == 0 {
			return false
		}
	} 
	return true
}

func cekKelengkapanPR(T tabRider, n int) bool {
/*{Mengembalikan nilai true apabila semua array waktu PR sudah terisi
dan nilai false apablia terdapat array waktu PR yang bernilai 0}*/
	var i int
	
	for i = 0; i < n; i++ {
		if T[i].PR == 0 {
			return false
		}
	} 
	return true
}