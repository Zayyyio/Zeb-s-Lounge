package model_test

import (
	"fmt"

	model "sekolahbeta/MiniProjek3/Model"
	"sekolahbeta/MiniProjek3/config"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestTambahBuku(t *testing.T) {
	Init()

	bookData := model.Perpus{
		Judul:   "Tensai Sakuragi",
		ISBN:    "0978",
		Penulis: "Yu Hasebe",
		Tahun:   2015,
		Stok:    5,
		Gambar:  "Wikipedia",
	}

	err := bookData.Tambah(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(bookData.ID)
}

func TestDeleteByID(t *testing.T) {
	Init()

	bookData := model.Perpus{
		Judul:   "Tensai Sakuragi",
		ISBN:    "0978",
		Penulis: "Yu Hasebe",
		Tahun:   2015,
		Stok:    5,
		Gambar:  "Wikipedia",
	}

	err := bookData.Tambah(config.Mysql.DB)
	assert.Nil(t, err)

	err = bookData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	Init()

	bookData := model.Perpus{
		Judul:   "Tensai Sakuragi",
		ISBN:    "0978",
		Penulis: "Yu Hasebe",
		Tahun:   2015,
		Stok:    5,
		Gambar:  "Wikipedia",
	}

	err := bookData.Tambah(config.Mysql.DB)
	assert.Nil(t, err)

	data, err := bookData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(data)

}

func TestGetAll(t *testing.T) {
	Init()

	bookData := model.Perpus{
		Judul:   "Oreno Jitsuokurou Mitsetearu",
		ISBN:    "0035",
		Penulis: "Kenji Fujima",
		Tahun:   2022,
		Stok:    17,
		Gambar:  "Pinterest",
	}

	err := bookData.Tambah(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := bookData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestUpdateByID(t *testing.T) {
	Init()

	booksData := model.Perpus{
		Judul:   "Tensai Sakuragi",
		ISBN:    "0978",
		Penulis: "Yu Hasebe",
		Tahun:   2015,
		Stok:    5,
		Gambar:  "Wikipedia",
	}

	err := booksData.Tambah(config.Mysql.DB)
	assert.Nil(t, err)

	booksData.Judul = "Test Update Aja"

	err = booksData.UpdateOne(config.Mysql.DB, 1)
	assert.Nil(t, err)

}
