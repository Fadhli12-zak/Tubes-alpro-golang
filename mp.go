package main

import "fmt"

const NMAX int = 100

type pegawai struct {
	nama     string
	jobdesk  string
	domisili string
	umur     int
	tugas    tugas
}
type tugas struct {
	proyek   string
	tenggat  string
	anggaran int
}
type tabInt [NMAX]pegawai

func main() {
	var A tabInt
	var x, y int
	var z string
	// var stop bool = true
	fmt.Println("=================================\n   Aplikasi Manajemen Proyek     \n          Created by:            \n       Khansa Aulia Fauziah      \n      Fadhli muhammad Dzaki      \n================================= ")
	menu := "Menu:\n1. Input Employee Data\n2. Search Employee Data\n3. Edit Employee Data\n4. Print Employee Data\n5.  Hapus data\n0. Exit\nMasukan Nomor Pilihan Anda: "
	fmt.Print(menu)
	fmt.Scan(&y)
	for y != 0 {
		if y == 1 {
			isi(&A, &x)
		}
		if y == 2 {
			fmt.Scan(&z)
			findSequential(A, x, z)
		}
		if y == 3 {
			fmt.Print("masukan nama yang ingin di rubah: ")
			fmt.Scan(&z)
			edit(&A, x, z)
		}
		if y == 4 {
			cetak(A, x)
		}
		if y == 5 {
			fmt.Println("hapus data:")
			fmt.Scan(&z)
			delete(&A, &x, z)
		}
		fmt.Print(menu)
		fmt.Scan(&y)
	}

}

func isi(P *tabInt, n *int) {
	var A tabInt
	if *n >= NMAX {
		return
	}
	for *n < NMAX {
		fmt.Print("Masukan nama anda (masukan x jika selesai): ")
		fmt.Scan(&A[*n].nama)
		if A[*n].nama == "x" {
			fmt.Println("Data selesai ditambahkan")
			return
		} else {
			fmt.Print("Masukan jobdesk anda: ")
			fmt.Scan(&A[*n].jobdesk)
			fmt.Print("Masukan domisili anda: ")
			fmt.Scan(&A[*n].domisili)
			fmt.Print("Masukan umur anda: ")
			fmt.Scan(&A[*n].umur)
			fmt.Print("Masukan proyek anda: ")
			fmt.Scan(&A[*n].tugas.proyek)
			fmt.Print("Masukan tenggat tugas anda: ")
			fmt.Scan(&A[*n].tugas.tenggat)
			fmt.Print("Masukan anggaran tugas anda: ")
			fmt.Scan(&A[*n].tugas.anggaran)
			P[*n] = A[*n]
			*n++
		}

		if A[*n].nama == A[*n+1].nama {
			fmt.Print("Nama tersebut sudah terdaftar")
		}
	}
}

func cetak(P tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur, P[i].tugas.proyek, P[i].tugas.tenggat, P[i].tugas.anggaran)
	}
}

func findSequential(P tabInt, n int, x string) int {
	var index int
	index = -1
	var status bool = true
	var i int = 0
	for i < n && status {
		if P[i].nama == x || P[i].jobdesk == x || P[i].tugas.proyek == x {
			status = false
			index = i
			fmt.Println(P[i].nama, P[i].jobdesk, P[i].domisili, P[i].umur, P[i].tugas.proyek, P[i].tugas.tenggat, P[i].tugas.anggaran)
		} else {
			fmt.Println("Data tidak ditemukan")
			break
		}
		i++
	}
	return index
}

// func findBinary() {
// 	left =
// }

func edit(P *tabInt, n int, x string) {
	var finds int
	var datInt int
	var datStr string
	finds = findSequential(*P, n, x)
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

		fmt.Print("Masukan Proyek: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].tugas.proyek = datStr
		}

		fmt.Print("Masukan Tenggat: ")
		fmt.Scan(&datStr)
		if datStr != "-" {
			P[finds].tugas.tenggat = datStr
		}

		fmt.Print("Masukan Anggaran: ")
		fmt.Scan(&datInt)
		if datInt != 0 {
			P[finds].tugas.anggaran = datInt
		}

	} else {
		fmt.Println("Data Tidak ditemukan")
	}

}
func delete(P *tabInt, n *int, x string) {
	var ketemu int

	ketemu = findSequential(*P, *n, x)
	if ketemu != -1 {
		for i := ketemu; i < *n-1; i++ {
			P[i] = P[i+1]
		}
		*n = *n - 1
	}
}

func rekap(P *tabInt, n *int, x string) {
	
}
