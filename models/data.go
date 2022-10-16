package models

import "time"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Password string
}

type As struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Slug        string
	ASN         uint
}

type Subnet struct {
	ID        uint `gorm:"primaryKey"`
	Address   string
	AsId      int
	As        As
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Target struct {
	ID          uint `gorm:"primaryKey"`
	Addr        string
	Description string
	SubnetId    uint
	Subnet      Subnet
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
