package main

import (
	"fmt"
)
var data [22]int
func main() {
	var menit, detik, miliDetik int
	for i := 0; i < 22; i++{
		fmt.Scan(&data[i])
	}
	for i := 0; i < 22; i++{
		menit = data[i] / 60000
		data[i]%= 60000
		detik = data[i] / 1000
		data[i]%= 1000
		miliDetik = data[i]
		fmt.Println(menit, detik, miliDetik)
	}

}

// 1197392
// 1206510
// 1213312
// 1206663
// 1194492
// 1200884
// 1202393
// 1205272
// 1201210
// 1206133
// 1189694
// 1217548
// 1191974
// 1193868
// 1203186
// 1198879
// 1201951
// 1208817
