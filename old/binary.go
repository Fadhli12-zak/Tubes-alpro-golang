package main

import "fmt"

type apalah struct {
	nama string
}

type info [100]apalah

func main() {
	var p info
	var N int
	// var x string
	Input(&p, &N)
	ctk(p, N)
	Srt(&p, N)
	ctk(p, N)
}
func Input(P *info, n *int) {
	var A info
	if *n >= 100 {
		return
	}
	for *n < 100 {
		fmt.Print("Masukan nama anda (masukan x jika selesai): ")
		fmt.Scan(&A[*n].nama)
		if A[*n].nama == "x" {
			fmt.Println("Data selesai ditambahkan")
			return
		}
	}
}

func Srt(P *info, n int) {
	var i, idx int
	var temp apalah
	i = 1
	for i < n {
		var j int
		idx = i - 1
		i = idx
		for j < n {
			if P[idx].nama > P[j].nama {
				idx = j
			}
			j++
		}
		temp = P[idx]
		P[idx] = P[i-1]
		P[i-1] = temp
		i = i + 1
	}
}

func ctk(P info, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(P[i].nama)
	}
}

// func binary(P info, n int, x string) int {
// 	var mid int
// 	var find int = -1
// 	var kr int = 0
// 	var kn int = n - 1
// 	for kr <= kn && find == -1 {
// 		mid = (kr + kn) / 2
// 		if x < P[mid].nama {
// 			kn = mid - 1
// 		} else if x > P[mid].nama {
// 			kr = mid + 1
// 		} else {
// 			find = mid
// 		}
// 	}
// 	fmt.Print(find)
// 	return find
// }
