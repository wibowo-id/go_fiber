package schema

type Wilayah struct {
	Kode string `json:"kode" gorm:"primary_key"`
	Nama string `json:"name"`
}
