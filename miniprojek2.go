package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/go-pdf/fpdf"
)

type Perpus struct {
	Judul         string `json:"judul"`
	JumlahHalaman int    `json:"jml_hal"`
	Genre         string `json:"genre"`
	KodeBuku      string `json:"kd_buku"`
	Pengarang     string `json:"pengarang"`
	Penerbit      string `json:"penerbit"`
	TahunTerbit   int    `json:"thn_terbit"`
}

var ListBuku []Perpus
var bukuKode = make(map[string]bool)

func tambahBuku() {
	judulUser := bufio.NewReader(os.Stdin)
	genreUser := bufio.NewReader(os.Stdin)
	pengarangUser := bufio.NewReader(os.Stdin)
	penerbitUser := bufio.NewReader(os.Stdin)

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
	bukuPelanggan, err := judulUser.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}
	bukuPelanggan = strings.Replace(bukuPelanggan, "\n", "", 1)

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
	for _, buku := range ListBuku {
		if buku.KodeBuku == kodePelanggan {
			fmt.Println("Kode Buku Sudah Terdaftar. Silahkan Gunakan Kode Lain.")
			return
		}
	}

	fmt.Print("silahkan Masukan Genre: ")
	genreBuku, err = genreUser.ReadString('\n')
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}
	genreBuku = strings.Replace(genreBuku, "\n", "", 1)

	fmt.Print("Siapa Pengarang Buku?:  ")
	pengarangBuku, err = pengarangUser.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}
	pengarangBuku = strings.Replace(pengarangBuku, "\n", "", 1)

	fmt.Print("Siapa Menerbitkan Buku itu?: ")
	penerbitBuku, err = penerbitUser.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi Error", err)
		return
	}
	penerbitBuku = strings.Replace(penerbitBuku, "\n", "", 1)

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
func simpanBuku(ch <-chan Perpus, wg *sync.WaitGroup, noPembaca int) {
	for perpus := range ch {
		dataJSON, err := json.Marshal(perpus)
		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile(fmt.Sprintf("Buku/%s.json", perpus.KodeBuku), dataJSON, 0644)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}
}

func lihatBuku(ch <-chan string, chBuku chan Perpus, wg *sync.WaitGroup) {
	var perpus Perpus
	for idBuku := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("Buku/%s", idBuku))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(dataJSON, &perpus)
		if err != nil {
			fmt.Println(err)
		}
		chBuku <- perpus
	}
	wg.Done()
}

func liatBuku() {

	fmt.Println("===================================")
	fmt.Println("Lihat Buku")
	fmt.Println("===================================")
	ListBuku = []Perpus{}

	listJsonBuku, err := os.ReadDir("Buku")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chBuku := make(chan Perpus, len(listJsonBuku))

	jumlahBuku := 5

	for i := 0; i < jumlahBuku; i++ {
		wg.Add(1)
		go lihatBuku(ch, chBuku, &wg)
	}

	for _, fileBuku := range listJsonBuku {
		fmt.Println(fileBuku.Name())
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(chBuku)

	for dataBuku := range chBuku {
		ListBuku = append(ListBuku, dataBuku)
	}

	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Judul: %s, Halaman: %d, kodeBuku: %s\n, Genre: %s\n, Pengarang: %s\n, Penerbit: %s\n, Tahun Terbit: %d\n",
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
	_, err := fmt.Scanln(&kodeHapusBuku)
	if err != nil {
		fmt.Println(err)
	}

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodeHapusBuku {
			// ListBuku = append(ListBuku[:i], append([]Perpus{ListBuku[i]}, ListBuku[i+1:]...)...)
			fmt.Println("Buku Berhasil Dihapus!")

			err := os.Remove(fmt.Sprintf("Buku/%s.json", buku.KodeBuku))
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("Buku tidak ditemukan.")
		}
		return

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
		editBukuToJson(ListBuku[i], fmt.Sprintf("buku/book-%s.json", buku.KodeBuku))

	}

}

func main() {
	os.Mkdir("Buku", 0777)

	pilihanUser := 0

	fmt.Println("===================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===================================")
	fmt.Println("Silahkan Pilih: ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Liat Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Ubah Buku")
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
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
		GeneratePdfBuku()
	case 6:
		os.Exit(0)
	}
	main()
}

func editBukuToJson(bukuEdit Perpus, fileName string) {
	encoded, err := json.MarshalIndent(bukuEdit, "", "    ")
	if err != nil {
		fmt.Println("Terjadi Error : ", err)

		return
	}

	err = ioutil.WriteFile(fileName, encoded, 0644)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)

		return
	}

}

func GeneratePdfBuku() {
	liatBuku()
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Pesanan ...")
	fmt.Println("=================================")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, perpus := range ListBuku {
		bukuText := fmt.Sprintf(
			"buku #%d:\nJudu; : %d\nJumlahHalaman : %d\nKodeBuku : %s\nGenre : %s\nPengarang : %s\nPenerbit : %d\nTahunTerbit",
			i+1, perpus.Judul, perpus.JumlahHalaman,
			perpus.KodeBuku, perpus.Genre, perpus.Pengarang,
			perpus.Penerbit, perpus.TahunTerbit)

		pdf.MultiCell(0, 10, bukuText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf"))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}
