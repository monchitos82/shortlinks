package models

type URL struct {
	ID     uint   `gorm:"primaryKey"`
	Handle string `gorm:"uniqueIndex"`
	URL    string
}
