package model

type User struct {
	Id             string `gorm:"unique;primaryKey;type:varchar(255)"`
	Username       string `gorm:"unique;type:varchar(100)"`
	HashedPassword string `gorm:"type:varchar(255)"`
	Role           string `gorm:"type:varchar(50)"`
}
