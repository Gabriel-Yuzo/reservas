package models

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email"`
}
