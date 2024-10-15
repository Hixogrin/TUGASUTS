package main

import (
	"fmt"
	"os"
)

type Buku struct {
	Nama   string
	Jumlah int
}

type Peminjaman struct {
	Username string
	Buku     string
	Jumlah   int
}

var bukuList = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

var peminjamanList []Peminjaman

// Username dan password mahasiswa
var userDB = map[string]string{
	"Denendro": "2406357942",
}

// Fungsi untuk login
func login() bool {
	var username, password string

	fmt.Println("=====================================================")
	fmt.Println("=      Selamat Datang di Perpustakaan Vokasi      =")
	fmt.Println("=====================================================")
	fmt.Println("            Ayo Menjadi Pembaca yang Baik!        ")
	fmt.Println("=====================================================")
	fmt.Print("Silahkan Input Username: ")
	fmt.Scanln(&username)

	fmt.Print("Silahkan Input Password: ")
	fmt.Scanln(&password)

	if pass, exists := userDB[username]; exists && pass == password {
		fmt.Println("=====================================================")
		fmt.Println("                Login Sukses!!!!                   ")
		fmt.Println("=====================================================")
		return true
	} else {
		fmt.Println("=====================================================")
		fmt.Println("         Login Gagal. Username atau Password salah! ")
		fmt.Println("=====================================================")
		return false
	}
}

// Fungsi untuk melihat informasi pengguna
func LihatInformasiPenggunaProgram() {
	fmt.Println("\n=====================================================")
	fmt.Println("---      Informasi Pengguna Program      ---")
	fmt.Println("=====================================================")
	fmt.Printf("Username: %s\n", "Denendro")
	fmt.Printf("NPM: %s\n", "2406357942")
	fmt.Printf("Jenis Kelamin: %s\n", "Laki-laki")
	fmt.Printf("Makanan Favorit: %s\n", "Nasi Padang")
	fmt.Printf("Minuman Favorit: %s\n", "Es Teh")
	fmt.Println("=====================================================")
}

// Fungsi untuk melihat daftar buku
func LihatDaftarBuku() {
	fmt.Println("\n=====================================================")
	fmt.Println("---            Daftar Buku            ---")
	fmt.Println("=====================================================")
	for _, buku := range bukuList {
		fmt.Printf("Nama Buku: %s, Jumlah: %d\n", buku.Nama, buku.Jumlah)
	}
	fmt.Println("=====================================================")
}

// Fungsi untuk menambah daftar buku
func TambahDaftarBuku() {
	var namaBuku string
	var jumlah int

	fmt.Println("\n=====================================================")
	fmt.Println("---          Tambah Daftar Buku          ---")
	fmt.Println("=====================================================")
	fmt.Print("Masukkan Nama Buku: ")
	fmt.Scanln(&namaBuku)

	fmt.Print("Masukkan Jumlah Buku: ")
	fmt.Scanln(&jumlah)

	bukuList = append(bukuList, Buku{Nama: namaBuku, Jumlah: jumlah})
	fmt.Println("=====================================================")
	fmt.Println("          Buku berhasil ditambahkan!           ")
	fmt.Println("=====================================================")
}

// Fungsi untuk menambah peminjaman buku
func TambahPeminjamanBuku() {
	var username string
	var jumlah int

	fmt.Println("\n=====================================================")
	fmt.Println("---         Tambah Peminjaman Buku         ---")
	fmt.Println("=====================================================")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)

	// Tampilkan daftar buku sebelum meminta input nama buku
	fmt.Println("\nDaftar Buku Tersedia untuk Dipinjam:")
	for i, buku := range bukuList {
		fmt.Printf("%d. %s (%d)\n", i+1, buku.Nama, buku.Jumlah)
	}

	var nomorBuku int
	fmt.Print("Masukkan Nomor Buku yang Ingin Dipinjam: ")
	fmt.Scanln(&nomorBuku)

	if nomorBuku < 1 || nomorBuku > len(bukuList) {
		fmt.Println("=====================================================")
		fmt.Println("          Nomor buku tidak valid!                 ")
		fmt.Println("=====================================================")
		return
	}

	namaBuku := bukuList[nomorBuku-1].Nama // Ambil nama buku dari daftar
	fmt.Print("Masukkan Jumlah Buku yang Dipinjam: ")
	fmt.Scanln(&jumlah)

	// Validasi jumlah buku
	if jumlah <= 0 || jumlah > bukuList[nomorBuku-1].Jumlah {
		fmt.Println("=====================================================")
		fmt.Println("      Jumlah buku tidak valid atau melebihi stok! ")
		fmt.Println("=====================================================")
		return
	}

	// Kurangi jumlah buku
	bukuList[nomorBuku-1].Jumlah -= jumlah
	peminjamanList = append(peminjamanList, Peminjaman{Username: username, Buku: namaBuku, Jumlah: jumlah})
	fmt.Println("=====================================================")
	fmt.Println("          Peminjaman berhasil dilakukan!          ")
	fmt.Println("=====================================================")
}

// Fungsi untuk melihat histori peminjaman buku
func HistoriPeminjamanBuku() {
	fmt.Println("\n=====================================================")
	fmt.Println("---         Histori Peminjaman Buku          ---")
	fmt.Println("=====================================================")
	if len(peminjamanList) == 0 {
		fmt.Println("                Belum ada histori peminjaman.     ")
	} else {
		for _, peminjaman := range peminjamanList {
			fmt.Printf("Username: %s, Buku: %s, Jumlah: %d\n", peminjaman.Username, peminjaman.Buku, peminjaman.Jumlah)
		}
	}
	fmt.Println("=====================================================")
}

// Fungsi untuk keluar dari program
func KeluarDariProgram() {
	fmt.Println("\n=====================================================")
	fmt.Println("     Terima kasih telah menggunakan program ini.     ")
	fmt.Println("=====================================================")
	os.Exit(0)
}

// Fungsi utama
func main() {
	if !login() {
		return
	}

	for {
		fmt.Println("\n=====================================================")
		fmt.Println("---                Menu Program                ---")
		fmt.Println("=====================================================")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar Dari Program")
		fmt.Print("Input Menu yang anda inginkan: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			LihatInformasiPenggunaProgram()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			KeluarDariProgram()
		default:
			fmt.Println("=====================================================")
			fmt.Println("          Pilihan tidak valid, coba lagi!          ")
			fmt.Println("=====================================================")
		}
	}
}