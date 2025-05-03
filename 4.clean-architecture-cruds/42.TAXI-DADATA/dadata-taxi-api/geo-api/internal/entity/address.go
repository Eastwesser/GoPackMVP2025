// internal/entity/address.go
package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID       uint   `gorm:"index"`
	RawAddress   string `gorm:"not null"`
	Result       string
	PostalCode   string
	Country      string
	Region       string
	CityArea     string
	CityDistrict string
	Street       string
	House        string
	GeoLat       float64
	GeoLon       float64
	QCGeo        int
}
