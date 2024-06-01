package main

import (
	"fmt"
	"time"
)

type Tanggal struct {
	Tgl  time.Time
	Data string
}

func main() {
	var data []Tanggal

	// Mengambil inputan data
	fmt.Print("Masukkan jumlah data: ")
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var tgl string
		var dataStr string

		fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
		fmt.Scanln(&tgl)

		fmt.Print("Masukkan data: ")
		fmt.Scanln(&dataStr)

		parsedTgl, err := time.Parse("2006-01-02", tgl)
		if err != nil {
			fmt.Println("Format tanggal tidak valid")
			continue
		}

		data = append(data, Tanggal{parsedTgl, dataStr})
	}

	// Selection Sort
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j].Tgl.Before(data[minIndex].Tgl) {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}

	fmt.Println("Data setelah diurutkan:")
	for _, d := range data {
		fmt.Println(d.Tgl.Format("2006-01-02"), d.Data)
	}
}
