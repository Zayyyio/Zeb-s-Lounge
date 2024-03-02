package main

import (
	"fmt"
	"os"
)

type Perpus struct {
	Judul         string
	JumlahHalaman int
	Genre         string
	KodeBuku      string
	Pengarang     string
	Penerbit      string
	TahunTerbit   int
}

var ListBuku []Perpus

func tambahBuku() {
	bukuPelanggan := ""
	genreBuku := ""
	pengarangBuku := ""
	penerbitBuku := ""
	kodePelanggan := ""
	thnTerbitBuku := 0
	halaman := 0

	fmt.Println("===================================")
	fmt.Println("           Tambah Buku             ")
	fmt.Println("===================================")
	fmt.Println("Silahkan Isi Judul Buku: ")
	_, err := fmt.Scanln(&bukuPelanggan)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("silahkan Masukan Halaman: ")
	_, err = fmt.Scanln(&halaman)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("silahkan Masukan Kode: ")
	_, err = fmt.Scanln(&kodePelanggan)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("silahkan Masukan Genre: ")
	_, err = fmt.Scanln(&genreBuku)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("Siapa Pengarang Buku?:  ")
	_, err = fmt.Scanln(&pengarangBuku)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("Siapa Menerbitkan Buku itu?: ")
	_, err = fmt.Scanln(&penerbitBuku)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	fmt.Print("Tahun Kapan Diterbitkan?: ")
	_, err = fmt.Scanln(&thnTerbitBuku)
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}

	ListBuku = append(ListBuku, Perpus{
		Judul:         bukuPelanggan,
		JumlahHalaman: halaman,
		KodeBuku:      kodePelanggan,
		Genre:         genreBuku,
		Pengarang:     pengarangBuku,
		Penerbit:      penerbitBuku,
		TahunTerbit:   thnTerbitBuku,
	})
}

func liatBuku() {
	fmt.Println("===================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===================================")
	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Judul: %s, Halaman: %d, kodeBuku: %s, Genre: %s, Pengarang: %s, Penerbit: %s, Tahun Terbit: %d\n",
			urutan+1,
			buku.Judul,
			buku.JumlahHalaman,
			buku.KodeBuku,
			buku.Genre,
			buku.Pengarang,
			buku.Penerbit,
			buku.TahunTerbit,
		)
	}
}

func hapusBuku() {
	kodeHapusBuku := ""
	fmt.Println("===================================")
	fmt.Println("Hapus Buku")
	fmt.Println("===================================")
	liatBuku()
	fmt.Println("===================================")
	fmt.Println("Masukkan Kode Buku yang ingin DiHapus")
	fmt.Scanln(&kodeHapusBuku)

	for i, buku := range ListBuku {
		if buku.KodeBuku == kodeHapusBuku {
			ListBuku = append(ListBuku[:i], append([]Perpus{ListBuku[i]}, ListBuku[i+1:]...)...)
			fmt.Println("Buku Berhasil Dihapus!")

		}

	}
}

func editBuku() {
	pergantianBuku := ""
	var kodeEdit string

	fmt.Println("===================================")
	fmt.Println("Ubah Buku")
	liatBuku()
	fmt.Println("===================================")
	fmt.Print("Masukkan kode buku yang akan diubah: ")
	fmt.Scanln(&kodeEdit)

	for i, buku := range ListBuku {
		if buku.KodeBuku == kodeEdit {
			fmt.Println("===================================")
			fmt.Print("Ingin Merubah yang mana?:  ")
			fmt.Scanln(&pergantianBuku)
		}
		if pergantianBuku == "Judul" {
			var judulBaru string
			fmt.Println("===================================")
			fmt.Println("Judul Lama : ", ListBuku[i].Judul)
			fmt.Println("Silahkan masukkan Judul baru Anda: ")
			fmt.Scanln(&judulBaru)
			ListBuku[i].Judul = judulBaru

		} else if pergantianBuku == "Halaman" {
			var halBaru int
			fmt.Println("===================================")
			fmt.Println("Halaman Lama : ", ListBuku[i].JumlahHalaman)
			fmt.Println("Silahkan masukkan Halaman baru Anda: ")
			fmt.Scanln(&halBaru)
			ListBuku[i].JumlahHalaman = halBaru

		} else if pergantianBuku == "Kode" || pergantianBuku == "kode" {
			fmt.Println("Maaf, Kode Buku tidak dapat dirubah.")

		} else if pergantianBuku == "Genre" {
			var genreBaru string

			fmt.Println("===================================")
			fmt.Println("Genre Lama : ", ListBuku[i].Genre)
			fmt.Println("Silahkan Masukkan Genre Baru Anda: ")
			fmt.Scanln(&genreBaru)
			ListBuku[i].Genre = genreBaru

		} else if pergantianBuku == "Pengarang" {
			var pengarangBaru string

			fmt.Println("===================================")
			fmt.Println("Pengarang Lama: ", ListBuku[i].Pengarang)
			fmt.Println("Silahkan Masukkan Pengarang Baru Anda")
			fmt.Scanln(&pengarangBaru)
			ListBuku[i].Pengarang = pengarangBaru

		} else if pergantianBuku == "Penerbit" {
			var penerbitBaru string
			fmt.Println("===================================")
			fmt.Println("Penerbit Lama : ", ListBuku[i].Penerbit)
			fmt.Println("Silahkan Masukkan Penerbit Baru Anda: ")
			fmt.Scanln(&penerbitBaru)
			ListBuku[i].Penerbit = penerbitBaru

		} else if pergantianBuku == "Tahun Terbit" {
			var thnTerbitBaru int

			fmt.Println("===================================")
			fmt.Println("Tahun Terbit Lama : ", ListBuku[i].TahunTerbit)
			fmt.Scanln(&thnTerbitBaru)
			ListBuku[i].TahunTerbit = thnTerbitBaru
		}

	}

}
func main() {
	pilihanUser := 0

	fmt.Println("===================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===================================")
	fmt.Println("Silahkan Pilih: ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Liat Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Ubah Buku")
	fmt.Println("5. Keluar")
	fmt.Println("===================================")
	fmt.Println("Masukkan Pilihanmu: ")
	_, err := fmt.Scanln(&pilihanUser)
	if err != nil {
		fmt.Println("terjadi Error", err)
	}

	switch pilihanUser {
	case 1:
		tambahBuku()
	case 2:
		liatBuku()
	case 3:
		hapusBuku()
	case 4:
		editBuku()
	case 5:
		os.Exit(0)
	}
	main()
}
