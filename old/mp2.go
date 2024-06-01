package main

import "fmt"

const NMAXp int = 100
const NMAXt int = 100

type pegawai struct {
	nama     string
	jobdesk  string
	domisili string
	umur     int
}
type proyek struct {
	tugas     string
	tenggat   string
	anggaran  int
	prioritas string
}
type tabInt [NMAXp]pegawai
type tgsInt [NMAXt]proyek

func main() {
	var B tgsInt
	var A tabInt
	var x, y, w, v int
	var z string
	// var stop bool = true
	fmt.Println("=================================\n   Aplikasi Manajemen Proyek     \n          Created by:            \n       Khansa Aulia Fauziah      \n      Fadhli muhammad Dzaki      \n================================= ")
	menu := "Menu:\n1. isi daftar tugas\n2. Isi data pegawai\n3. cari Data pegawai\n4. edit data pegawai\n5. cetak data pegawai\n6. cetak data tugas\n7. Hapus data pegawai\n8. Urutkan prioritas\n0. Exit\nMasukan Nomor Pilihan Anda: "
	fmt.Print(menu)
	fmt.Scan(&y)
	for y != 0 {
		if y == 1 {
			isiTugas(&B, &w)
		}
		if y == 2 {
			isiNama(&A, &x)
		}
		if y == 3 {
			fmt.Println("cari berdasarkan:\n1.Nama\n2.Proyek\n3.jobdesk")
			fmt.Scan(&v)
			if v == 1 {
				fmt.Scan(&z)
				findN(A, x, z)
			}
			if v == 2 {
				fmt.Scan(&z)
				findP(B, w, z)
			}
			// if v == 3 {
			// 	fmt.Scan(&z)
			// 	findJdk(A, x, z)
			// }
		}
		if y == 4 {
			fmt.Println("pilih data yang ingin anda edit\n1. nama\n2. tugas")
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				fmt.Println("masukan nama yang ingin di rubah:")
				fmt.Scan(&z)
				editNama(&A, x, z)
			}
			if pilih == 2 {
				fmt.Println("Masukan tugas yang ingin di edit")
				fmt.Scan(&z)
				editTugas(&B, w, z)
			}
		}
		if y == 5 {
			cetak(A, x)
		}
		if y == 6 {
			cetakTugas(B, w)
		}
		if y == 7 {
			fmt.Println("pilih data yang ingin anda hapus\n1. nama\n2. tugas")
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				fmt.Println("hapus data:")
				fmt.Scan(&z)
				deleteNama(&A, &x, z)
			}
			if pilih == 2 {
				fmt.Println("hapus data:")
				fmt.Scan(&z)
				deleteTugas(&B, &w, z)
			}
		}
		if y == 8 {
			fmt.Println("urutan data secara ascending:")
			selSrt(&B, w)
			cetakTugas(B, w)
		}
		fmt.Print(menu)
		fmt.Scan(&y)
	}
}

func isiTugas(P *tgsInt, q *int) {
	var A tgsInt
	var berhenti bool = true
	if *q >= NMAXt && berhenti {
		berhenti = false
	}
	for *q < NMAXt && berhenti {
		fmt.Print("Masukan proyek anda(masukan x jika selesai) ")
		fmt.Scan(&A[*q].tugas)
		if A[*q].tugas == "x" {
			fmt.Println("Data selesai ditambhkan")
			berhenti = false
		} else {
			fmt.Print("Masukan tenggat tugas anda: ")
			fmt.Scan(&A[*q].tenggat)
			fmt.Print("Masukan anggaran tugas anda: ")
			fmt.Scan(&A[*q].anggaran)
			fmt.Print("Tipe prioritas tugas:")
			fmt.Scan(&A[*q].prioritas)
			P[*q] = A[*q]
			*q++
		}
	}
}

func isiNama(P *tabInt, n *int) {
	var A tabInt
	var berhenti bool = true
	if *n >= NMAXp && berhenti {
		berhenti = false
	}
	for *n < NMAXp && berhenti {
		fmt.Print("Masukan nama anda (masukan x jika selesai): ")
		fmt.Scan(&A[*n].nama)
		if A[*n].nama == "x" {
			fmt.Println("Data selesai ditambahkan")
			berhenti = false
		} else {
			fmt.Print("Masukan jobdesk anda: ")
			fmt.Scan(&A[*n].jobdesk)
			fmt.Print("Masukan domisili anda: ")
			fmt.Scan(&A[*n].domisili)
			fmt.Print("Masukan umur anda: ")
			fmt.Scan(&A[*n].umur)
			P[*n] = A[*n]
			*n++
		}
	}
}

func cetak(P tabInt, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(i+1, P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
	}
}
func cetakTugas(P tgsInt, q int) {
	var i int
	fmt.Println("tugas", "tenggat", "anggaran", "Prioritas")
	for i = 0; i < q; i++ {
		fmt.Println(i+1, P[i].tugas, P[i].tenggat, P[i].anggaran, P[i].prioritas)
	}
}

func findN(P tabInt, n int, x string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < n && status {
		if P[i].nama == x {
			status = false
			index = i
			fmt.Println(P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
		} else {
			status = false
			index = -1
			fmt.Print("Data tidak ditemukan")
		}
		i++
	}
	fmt.Print(index)
	return index
}
func findP(P tgsInt, q int, m string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < q && status {
		if P[i].tugas == m {
			status = false
			index = i
			fmt.Println(P[i].tugas, P[i].tenggat, P[i].anggaran, P[i].prioritas)
		} else {
			status = false
			index = -1
			fmt.Print("Data tidak ditemukan")
		}
		i++
	}
	return index
}

func findJdk(P tabInt, n int, x string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < n && status {
		if P[i].jobdesk == x {
			status = false
			index = i
			fmt.Println(P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur)
		}
		i++
	}
	return index
}
func editNama(P *tabInt, n int, x string) {
	var finds int
	var datInt int
	var datStr string
	finds = findN(*P, n, x)
	if finds != -1 {
		fmt.Print("Msukan nama: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].nama = datStr
		}
		fmt.Print("Masukan jobdesk: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].jobdesk = datStr
		}

		fmt.Print("Masukan Domisili: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].domisili = datStr
		}

		fmt.Print("Masukan umur: ")
		fmt.Scan(&datInt)
		if datInt != 0 {
			P[finds].umur = datInt
		}

	} else {
		fmt.Println("Data Tidak ditemukan")
	}
}
func editTugas(P *tgsInt, q int, m string) {
	var ketemu int
	var integer int
	var stringdata string
	ketemu = findP(*P, q, m)
	if ketemu != -1 {
		fmt.Print("Msukan nama: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].tugas = stringdata
		}
		fmt.Print("Masukan jobdesk: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].tenggat = stringdata
		}

		fmt.Print("Masukan Domisili: ")
		fmt.Scan(&stringdata)
		if stringdata != "-" {
			P[ketemu].prioritas = stringdata
		}

		fmt.Print("Masukan umur: ")
		fmt.Scan(&integer)
		if integer != 0 {
			P[ketemu].anggaran = integer
		}
	} else {
		fmt.Println("Data Tidak ditemukan")
	}
}

func deleteNama(P *tabInt, n *int, x string) {
	var ketemu int

	ketemu = findN(*P, *n, x)
	if ketemu != -1 {
		for i := ketemu; i < *n-1; i++ {
			P[i] = P[i+1]
		}
		*n = *n - 1
	}
	if ketemu == -1 {
		fmt.Println("data tidak ditemukan")
	}
}

func deleteTugas(P *tgsInt, q *int, m string) {
	var found int
	found = findP(*P, *q, m)
	if found == 0 {
		for i := found; i < *q-1; i++ {
			P[i] = P[i+1]
		}
		*q--
	}
	if found == -1 {
		fmt.Println("data tidak ditemukan")
	}
}
func selSrt(P *tgsInt, q int) {
	var i, idx int
	var temp proyek
	if q > NMAXt {
		q = NMAXt
	}
	i = 1
	for i < q {
		var j int
		idx = i - 1
		j = i
		for j < q {
			if P[idx].prioritas < P[j].prioritas {
				idx = j
			}
			j++
		}
		temp = P[idx]
		P[idx] = P[i-1]
		P[i-1] = temp
		i++
	}
}

// func binary(){

// }
