package model

type Note struct {
	Id          string `gorm:"primaryKey;type:varchar(255)"`
	UserId      string `gorm:"type:varchar(255)"`
	User        User
	Title       string `gorm:"type:varchar(100)"`
	Description string
	ImageUrl    string `gorm:"type:varchar(255)"`
}
