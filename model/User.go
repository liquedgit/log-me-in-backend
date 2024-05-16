package model

type User struct {
	Id             string `gorm:"unique;primaryKey"`
	Username       string `gorm:"unique"`
	HashedPassword string
	Role           string
}
