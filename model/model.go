package model

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup(){
	dsn := "host=127.0.0.1 user=postgres password=123 dbname=go-bitly-clone port=5432"
}
