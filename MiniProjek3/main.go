package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	model "sekolahbeta/MiniProjek3/Model"
	"sekolahbeta/MiniProjek3/config"
	"strconv"

	"time"

	"strings"

	"github.com/joho/godotenv"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ListBuku []model.Perpus

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load env file, using system env")
	}

}

func main() {
	InitEnv()
	config.OpenDB()

	os.Mkdir("pdf", 0777)

	pilihanUser := 0

	fmt.Println("|==============================================|")
	fmt.Println("|Aplikasi Manajemen Daftar Buku Perpustakaan   |")
	fmt.Println("|==============================================|")
	fmt.Println("Silahkan Pilih: ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Liat Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Ubah Buku")
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
	fmt.Println("=================================================")
	fmt.Println("Masukkan Pilihanmu: ")
	_, err := fmt.Scanln(&pilihanUser)
	if err != nil {
		fmt.Println("terjadi Error", err)
	}

	switch pilihanUser {
	case 1:
		TambahBuku()
	case 2:
		ListBukuDB()
	case 3:
		HapusBuku(&model.Model{})
	case 4:
		EditBuku(&model.Perpus{})
	case 5:
		PrintBuku()
	case 6:
		os.Exit(0)
	}
	main()

}

func ImportCsvFile(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return err
	}
	book := []model.Perpus{}

	for index, book := range records {

		if index == 0 {
			continue
		}
		IDreq, err := strconv.Atoi(book[0])
		if err != nil {
			IDreq = 0
		}
		stokReq, err := strconv.Atoi(book[6])
		if err != nil {
			stokReq = 0
		}

		tahunReq, err := strconv.Atoi(book[3])
		if err != nil {
			tahunReq = 0
		}

		book = append(book, model.Perpus{
			ISBN:    book[1],
			Penulis: book[2],
			Tahun:   uint(tahunReq),
			Judul:   book[4],
			Gambar:  book[5],
			Stok:    uint(stokReq),
			Model: model.Model{
				ID: IDreq,
			},
		})

	}

	for _, books := range book {
		books.CreatedAt = time.Now()
		books.UpdatedAt = time.Now()

		res, err := books.GetBySpecific(config.Mysql.DB)
		if err != nil {
			if err.Error() != "record not found" {
				return err

			} else {

				err = books.Tambah(config.Mysql.DB)
				if err != nil {
					return fmt.Errorf("failed to import data, err : %s", err.Error())

				}
			}
		} else {
			books.ID = res.ID
			books.CreatedAt = res.CreatedAt

			err = books.UpdateOne(config.Mysql.DB, 3)
			if err != nil {
				return fmt.Errorf("failed to update data, error %s", err.Error())
			}

		}
	}
	return nil
}

func ListBukuDB() {
	db := config.OpenDB()

	books := BookDB(db)
	for _, buku := range books {

		fmt.Printf("| Judul : %s, ID : %d\n", "| ISBN Buku : %s\n", "| Tahun : %s\n",
			"| - Penulis : %s\n", "| - Stok : %d\n", "| - Gambar : %s\n",

			buku.Judul,
			buku.ID,
			buku.ISBN,
			buku.Tahun,
			buku.Penulis,
			buku.Stok,
			buku.Gambar,
		)

		fmt.Println("-----------------------------------------------")
	}

}

func BookDB(db *gorm.DB) []model.Perpus {
	var books []model.Perpus
	db.Limit(20).Find(&books)
	return books
}

func TambahBuku() {
	userInput := bufio.NewReader(os.Stdin)

	isbnTambah := ""
	judulBukuTambah := ""
	penulisBukuTambah := ""
	tahunTerbitBukuTambah := 0
	stokBukuTambah := 0
	gambarBukuTambah := ""

	fmt.Println("|=============================================|")
	fmt.Println("|               Tambah Buku Baru              |")
	fmt.Println("|=============================================|")
	fmt.Print("| Silahkan Isi ISBN Buku : ")
	isbnTambah, _ = userInput.ReadString('\n')
	isbnTambah = strings.TrimSpace(isbnTambah)

	for _, buku := range ListBuku {
		if buku.ISBN == isbnTambah {
			fmt.Println("|=============================================|")
			fmt.Println("|          Maaf ISBN Buku Sudah Ada!          |")
			fmt.Println("|=============================================|")
			return
		}
	}

	fmt.Print("| Silahkan Isi Judul Buku : ")
	judulBukuTambah, _ = userInput.ReadString('\n')
	judulBukuTambah = strings.TrimSpace(judulBukuTambah)

	fmt.Print("| Silahkan Isi Penulis Buku : ")
	penulisBukuTambah, _ = userInput.ReadString('\n')
	penulisBukuTambah = strings.TrimSpace(penulisBukuTambah)

	fmt.Print("| Silahkan Isi Tahun Terbit Buku : ")
	_, err := fmt.Scanln(&stokBukuTambah)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	fmt.Print("| Silahkan Isi Stok Buku : ")
	_, err = fmt.Scanln(&stokBukuTambah)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	fmt.Print("| Silahkan Isi Link/Sumber Gambar Buku : ")
	gambarBukuTambah, _ = userInput.ReadString('\n')
	gambarBukuTambah = strings.TrimSpace(gambarBukuTambah)

	booksData := model.Perpus{
		Judul:   judulBukuTambah,
		ISBN:    isbnTambah,
		Penulis: penulisBukuTambah,
		Tahun:   uint(tahunTerbitBukuTambah),
		Stok:    uint(stokBukuTambah),
		Gambar:  gambarBukuTambah,
	}

	err = booksData.Tambah(config.Mysql.DB)
	if err != nil {
		fmt.Println("Terjadi Error")
	}
	fmt.Println("| Berikut adalah ID Buku : ", booksData.ID)

	fmt.Println("|=============================================|")
	fmt.Println("|            Buku Berhasil Ditambahkan!       |")
	fmt.Println("|=============================================!")
	var inputBaru string
	fmt.Print("| Apakah Anda ingin menambah buku lagi? (y/n)?: ")
	_, err = fmt.Scanln(&inputBaru)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	if strings.ToLower(inputBaru) == "y" {
		TambahBuku()
	} else {
		main()
	}

}

func HapusBuku(bk *model.Model) {

	fmt.Println("+=============================================+")
	fmt.Println("|                 Hapus Buku                  |")
	fmt.Println("+=============================================+")
	ListBukuDB()
	fmt.Println("+=============================================+")
	var kodeHapusBuku uint
	fmt.Print("| Masukkan ID Buku yang akan dihapus: ")
	_, err := fmt.Scanln(&kodeHapusBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	booksData := model.Perpus{
		Model: model.Model{
			ID: kodeHapusBuku,
		},
	}

	err = booksData.DeleteByID(config.Mysql.DB)
	if err != nil {
		fmt.Println("Terjadi Error : ", err.Error())
	}

}

func EditBuku(buku *model.Perpus) {
	var pergantianBuku = ""

	fmt.Println("|=============================================|")
	fmt.Println("|                  Edit Buku                  |")
	fmt.Println("|=============================================|")
	ListBukuDB()
	fmt.Println("|=============================================|")
	var IDEditBuku string
	fmt.Print("| Masukkan ID Buku yang akan diedit: ")
	_, err := fmt.Scanln(&IDEditBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	for i, buku := range ListBuku {
		if buku.ISBN == IDEditBuku {
			fmt.Println("===================================")
			fmt.Print("Ingin Merubah yang mana?:  ")
			fmt.Scanln(&pergantianBuku)
		}
		if pergantianBuku == "Judul" {
			userInput := bufio.NewReader(os.Stdin)

			fmt.Println("===================================")
			fmt.Println("Judul Lama : ", ListBuku[i].Judul)
			fmt.Println("Silahkan masukkan Judul baru Anda: ")
			judulBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			judulBukuBaru = strings.Replace(judulBukuBaru, "\n", "", 1)

		} else if pergantianBuku == "ISBN" {
			fmt.Println("===================================")
			fmt.Println("ISBN Lama : ", ListBuku[i].ISBN)
			fmt.Println("Silahkan masukkan ISBN baru Anda: ")

			_, err = fmt.Scanln(&buku.ISBN)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

		} else if pergantianBuku == "ID" || pergantianBuku == "id" {
			fmt.Println("Maaf, ID Buku tidak dapat dirubah.")

			var inputUlang string
			fmt.Print("| Apakah Anda ingin mengulangi lagi? (y/n)?: ")
			_, err = fmt.Scanln(&inputUlang)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			if strings.ToLower(inputUlang) == "y" {
				EditBuku(&model.Perpus{})

			} else {

				main()
			}

		} else if pergantianBuku == "Tahun" {
			userInput := bufio.NewReader(os.Stdin)

			fmt.Println("===================================")
			fmt.Println("Tahun Lama: ", ListBuku[i].Tahun)

			fmt.Println("Silahkan Masukkan Tahun Baru Anda")
			TahunBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			TahunBukuBaru = strings.Replace(TahunBukuBaru, "\n", "", 1)

		} else if pergantianBuku == "Penulis" {
			userInput := bufio.NewReader(os.Stdin)

			fmt.Println("===================================")
			fmt.Println("Penulis Lama : ", ListBuku[i].Penulis)
			fmt.Println("Silahkan Masukkan Penulis Baru Anda: ")

			PenulisBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			PenulisBukuBaru = strings.Replace(PenulisBukuBaru, "\n", "", 1)

		} else if pergantianBuku == "Stok" {

			fmt.Println("===================================")
			fmt.Println("Stok Terbit Lama : ", ListBuku[i].Stok)
			_, err = fmt.Scanln(&buku.Stok)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return

			}

		} else if pergantianBuku == "Gambar" {
			userInput := bufio.NewReader(os.Stdin)

			fmt.Println("===================================")
			fmt.Println("Gambar Lama : ", ListBuku[i].Gambar)

			GambarBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			GambarBukuBaru = strings.Replace(GambarBukuBaru, "\n", "", 1)

		}

		if pergantianBuku == "Semua" {
			userInput := bufio.NewReader(os.Stdin)

			fmt.Println("===================================")
			fmt.Println("Judul Lama : ", ListBuku[i].Judul)
			fmt.Println("Silahkan masukkan Judul baru Anda: ")
			judulBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			judulBukuBaru = strings.Replace(judulBukuBaru, "\n", "", 1)

			fmt.Println("===================================")
			fmt.Println("ISBN Lama : ", ListBuku[i].ISBN)
			fmt.Println("Silahkan masukkan ISBN baru Anda: ")

			_, err = fmt.Scanln(&buku.ISBN)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}
			fmt.Println("===================================")
			fmt.Println("Tahun Lama: ", ListBuku[i].Tahun)

			fmt.Println("Silahkan Masukkan Tahun Baru Anda")
			TahunBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			TahunBukuBaru = strings.Replace(TahunBukuBaru, "\n", "", 1)

			fmt.Println("===================================")
			fmt.Println("Penulis Lama : ", ListBuku[i].Penulis)
			fmt.Println("Silahkan Masukkan Penulis Baru Anda: ")

			PenulisBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			PenulisBukuBaru = strings.Replace(PenulisBukuBaru, "\n", "", 1)

			fmt.Println("===================================")
			fmt.Println("Stok Terbit Lama : ", ListBuku[i].Stok)
			_, err = fmt.Scanln(&buku.Stok)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return

			}
			fmt.Println("===================================")
			fmt.Println("Gambar Lama : ", ListBuku[i].Gambar)

			GambarBukuBaru, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			GambarBukuBaru = strings.Replace(GambarBukuBaru, "\n", "", 1)

		}

	}

}

func PrintBuku() {

	fmt.Println("|=============================================|")
	fmt.Println("|                 Print Buku                  |")
	fmt.Println("|=============================================|")
	ListBukuDB()
	fmt.Println("|=============================================|")

	var IDPrintBuku int
	fmt.Print("| Masukkan ID Buku yang ingin Diprint: ")
	_, err := fmt.Scanln(&IDPrintBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	buku := model.Perpus{
		Model: model.Model{
			ID: uint(IDPrintBuku),
		},
	}

	data, err := buku.GetByID(config.Mysql.DB)
	if err != nil {
		fmt.Println("Terjadi Error : ", err.Error())
	}

	GeneratePDF([]model.Perpus{data})
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=   Buku Berhasil Di-Print!   =+=+=+= |")
	fmt.Println("+=============================================+")
}

func GeneratePDF(bk []model.Perpus) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times New Roman", "B", 18)
	pdf.Cell(40, 10, "BUKU")

	for _, buku := range bk {
		pdf.Ln(10)

		pdf.SetFont("Times New Roman", "", 14)
		pdf.Cell(40, 10, fmt.Sprintf("ISBN : %s", buku.ISBN))
		pdf.Ln(5)
		pdf.Cell(40, 10, fmt.Sprintf("Judul : %s", buku.Judul))
		pdf.Ln(5)
		pdf.Cell(40, 10, fmt.Sprintf("Penulis : %s", buku.Penulis))
		pdf.Ln(5)
		pdf.Cell(40, 10, fmt.Sprintf("Tahun : %d", buku.Tahun))
		pdf.Ln(5)
		pdf.Cell(40, 10, fmt.Sprintf("Gambar : %s", buku.Gambar))
		pdf.Ln(5)
		pdf.Cell(40, 10, fmt.Sprintf("Stok : %d", buku.Stok))
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(fmt.Sprintf("pdf/book-%d.pdf", time.Now().Unix()))
	if err != nil {
		fmt.Println("Terjadi Error :", err)
		return
	}

}
