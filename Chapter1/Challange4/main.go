package main

import (
	"fmt"
	"os"
	"strconv"
)

// struct untuk menyimpan data teman
type Teman struct {
    Absen    int
    Nama     string
    Alamat   string
    Pekerjaan string
    Alasan   string
}

// function untuk mengambil data teman berdasarkan nomor absen
func getTeman(absen int) *Teman {
    // data teman-teman kelas
    teman1 := &Teman{1, "Mamat", "Jakarta", "Developer", "Ingin mempelajari bahasa pemrograman baru"}
    teman2 := &Teman{2, "Asep", "Bandung", "Designer", "Tertarik dengan fitur-fitur Golang"}
    teman3 := &Teman{3, "Bambang", "Surabaya", "Project Manager", "Mencari bahasa pemrograman yang cepat dan efisien"}

    // mencari data teman berdasarkan nomor absen
    switch absen {
    case 1:
        return teman1
    case 2:
        return teman2
    case 3:
        return teman3
    default:
        return nil
    }
}

func main() {
    // mengambil nomor absen dari argumen terminal
    absen, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Nomor absen harus berupa angka")
        return
    }

    // mengambil data teman berdasarkan nomor absen
    teman := getTeman(absen)
    if teman == nil {
        fmt.Println("Teman tidak ditemukan")
        return
    }

    // menampilkan data teman
    fmt.Println("Data teman dengan nomor absen", teman.Absen)
    fmt.Println("Nama\t\t:", teman.Nama)
    fmt.Println("Alamat\t\t:", teman.Alamat)
    fmt.Println("Pekerjaan\t:", teman.Pekerjaan)
    fmt.Println("Alasan\t\t:", teman.Alasan)
}