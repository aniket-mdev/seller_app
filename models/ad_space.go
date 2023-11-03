package models

import (
	"time"

	"gorm.io/gorm"
)

type CreateAdSpaceDTO struct {
	Description string `json:"description" validate:"required"`
	BasePrice   int64  `json:"base_price" validate:"required"`
	Status      string `json:"status" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
}

type AdSpace struct {
	gorm.Model
	Description string `json:"description"`
	BasePrice   int64  `json:"base_price"`
	Status      string `json:"status"`
	EndTime     string `json:"end_time"`
}

// create Ad Space
func CreateAdSpace(db *gorm.DB, AdSpace *AdSpace) (err error) {
	err = db.Create(AdSpace).Error
	return
}

// get ad space
func GetAdSpaces(db *gorm.DB, adSpace *[]AdSpace) (err error) {
	err = db.Find(&adSpace).Error
	return
}

// get ad by id
func GetAdSpace(db *gorm.DB, adSpace *AdSpace, id int) (err error) {
	err = db.Where("id = ?", id).First(adSpace).Error
	return
}

// update ad space
func UpdateAdSpace(db *gorm.DB, adSpace *AdSpace) (err error) {
	err = db.Save(adSpace).Error
	return
}

// delete adspace
func DeleteAdSpace(db *gorm.DB, adSpace *AdSpace, id int) (err error) {
	err = db.Where("id = ?", id).Delete(adSpace).Error
	return
}

func CheckAdSpaceEndTime(db *gorm.DB, adSpaces *[]AdSpace) (err error) {

	err = db.Where("str_to_date(?, '%Y-%m-%d %H:%i') > str_to_date(end_time, '%Y-%m-%d %H:%i') and status='open'", time.Now()).Find(adSpaces).Error
	return
}
