package models

type Logs struct {
	ID      int `gorm:"primaryKey"`
	EVENT   string
	OUTCOME string
	CLIENT  string
}
