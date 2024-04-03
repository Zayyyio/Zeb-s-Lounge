package model

import (
	"gorm.io/gorm"
)

type Perpus struct {
	Model

	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

func (buk *Perpus) Tambah(db *gorm.DB) error {

	err := db.
		Model(Perpus{}).
		Create(&buk).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (buk *Perpus) GetByID(db *gorm.DB) (Perpus, error) {
	res := Perpus{}

	err := db.
		Model(Perpus{}).
		Where("id =?", buk.ID).
		Take(&res).
		Error

	if err != nil {
		return Perpus{}, err
	}
	return res, nil
}

func (buk *Perpus) GetAll(db *gorm.DB) ([]Perpus, error) {
	res := []Perpus{}

	err := db.
		Model(Perpus{}).
		Find(&res).
		Error

	if err != nil {
		return []Perpus{}, err
	}
	return res, nil
}

func (buk *Perpus) UpdateOne(db *gorm.DB, ID uint) error {
	err := db.
		Model(Perpus{}).
		Select("ISBN", "Penulis", "Tahun", "Judul",
			"Gambar", "Stok").
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"ISBN":    buk.ISBN,
			"Penulis": buk.Penulis,
			"Tahun":   buk.Judul,
			"Gambar":  buk.Gambar,
			"Stok":    buk.Stok,
		}).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (buk *Perpus) DeleteByID(db *gorm.DB) error {

	if err := db.
		Model(Perpus{}).
		Where("id = ?", buk.Model.ID).
		First(&Perpus{}).
		Error; err != nil {
		return err
	}

	if err := db.Model(Perpus{}).
		Where("id = ?", buk.Model.ID).
		Delete(&Perpus{}).
		Error; err != nil {
		return err

	}

	return nil
}

func (bk *Perpus) GetBySpecific(db *gorm.DB) (Perpus, error) {
	res := Perpus{}

	query := db.Model(Perpus{})


	err := query.Take(&res).Error
	if err != nil {
		return Perpus{}, err
	}

	return res, nil
}
