package main
import "fmt"

type rider struct {
	name, no, nat, team                                                 string
    timeFP1, timePR, timeFP2, timeQ1, timeQ2, timeSPR, timeWUP, timeRAC int
    poinSPR, poinRAC, totalPoin                                         int
}

type subRider struct {
	name, no, nat, team, timeGap string
    time                         int
}

const NMAX int = 24

type tabRider [NMAX]rider

type tabSubRider [NMAX]subRider



func main() {
	var dataRider tabRider
	var dataFP1, dataPR, dataFP2, dataQ1, dataQ2, dataSPR, dataWUP, dataRAC tabSubRider
	var nDataRider int


}

func konversiWaktuTampilan(T  tabSubRider, posisi  integer) string {
/*mengembalikan string dengan format “(menit)’(detik).(miliDetik)s” jika posisi 1. Jika posisi > 1, maka mengembalikan string jarak dengan posisi pertama dengan . Jika waktu negatif, maka mengembalikan perolehan laps terakhir*/
	var menit, detik, miliDetik int

	T[]
	return
}

func hitungTotalPoin(S rider) int {
/*mengembalikan nilai total poin S rider*/
    return S.poinSPR + S.poinRAC
}

func inputData(T *tabRider, n *int) {
/*I.S. terdefinisi array T sebanyak n
 F.S. T berisi nilai. Banyak elemen (n) berisi nilai NMAX, jika n > NMAX*/
	var i int
    if *n > NMAX {
		*n = NMAX
	}

    for i = 0; i < *n; i++ {
		fmt.Scan(&T[i].name, &T[i].no, &T[i].nat, &T[i].team, &T[i].timeFP1, &T[i].timePR, &T[i].timeFP2, &T[i].timeQ1, &T[i].timeQ2, &T[i].timeSPR, &T[i].timeWUP, &T[i].timeRAC )
    }
}
// func inputRider(in/out T : tabRider, n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai. Banyak elemen (n) berisi nilai NMAX, jika n > NMAX*/

// func inputFP1(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeFP1*/

// func inputPR(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timePR*/

// func inputFP2(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeFP2*/

// func inputQ1(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeQ1*/

// func inputQ2(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeQ2*/

// func inputWUP(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeWUP*/

// func inputRAC(in/out T : tabRider, in n : integer)
// /*I.S. terdefinisi array T sebanyak n
//  F.S. T berisi nilai memasukkan field timeRAC*/

func urutWaktu(T *tabSubRider) {
/*I.S. terdefinisi array acak T1 sebanyak n
F.S. T berisi nilai terurut berdasarkan waktu*/
	var i, minPos, pass int
	var temp subPemain

	for pass = 1; pass < 6; pass++ {
		minPos = pass - 1
		for i = pass; i < 6; i++ {
			if T[i].time < T[minPos].time {
				minPos = i
			}
		}
		temp = T[minPos]
		T[minPos] = T[pass - 1]
		T[pass - 1] = temp
	}
}

func cetakTabelPemain(T tabRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team*/

}

func cetakTabelWaktu(T tabSubRider, n int) {
/*I.S. T array dan banyak elemen n terdefinisi
F.S. tercetak tabel dari field : name, no, nat, team, time*/

}

// func cetakTabelPoin(in T : tabRider, n int)
// /*I.S T array dan banyak elemen n terdefinisi
//  F.S tercetak tabel dari field : name, no, nat, team, totalPoin*/
