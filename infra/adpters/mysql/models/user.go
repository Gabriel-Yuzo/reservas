package models

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Username string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
}
