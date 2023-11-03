package models

import (
	"gorm.io/gorm"
)

type CreateBidderDTO struct {
	Name      string `json:"name" validate:"required"`
	BidAmount int64  `json:"bid_amount" validate:"required"`
	AdSpaceId int64  `json:"ad_space_id" validate:"required"`
}

type UpdateBiddeDTO struct {
	BidAmount int64 `json:"bid_amount" validate:"required"`
}

type Bidders struct {
	gorm.Model
	Name      string `json:"name"`
	BidAmount int64  `json:"bid_amount"`
	AdSpaceId int64  `json:"ad_space_id"`
	Status    string `json:"status"` // it will have a two status Open and Close, will update from cron where ad space end time
}

// create bidder for Ad-Space
func CreateBidde(db *gorm.DB, Bidders *Bidders) (err error) {
	err = db.Create(Bidders).Error
	return
}

// get bidders
func GetBidders(db *gorm.DB, bidder *[]Bidders) (err error) {
	err = db.Find(&bidder).Error
	return
}

// get bidders by ad space id
func GetBiddersByAdSpace(db *gorm.DB, bidders *[]Bidders, id int) (err error) {
	err = db.Where("ad_space_id = ?", id).Find(bidders).Error
	if err != nil {
		return
	}
	if len(*bidders) == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

// get bidder by name
func GetBidder(db *gorm.DB, bidders *[]Bidders, name string) (err error) {
	err = db.Where("name=?", name).Find(&bidders).Error
	if err != nil {
		return err
	}

	if len(*bidders) == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

// get bidde by bidde id
func GetBiddeById(db *gorm.DB, bidde *Bidders, id int) (err error) {
	err = db.Where("id = ?", id).First(&bidde).Error
	return
}

// update bidding
func UpdateBidder(db *gorm.DB, bidder *Bidders) (err error) {
	err = db.Save(bidder).Error
	return
}

// delete adspace
func DeleteBidder(db *gorm.DB, bidder *Bidders, id int) (err error) {
	err = db.Where("id = ?", id).Delete(bidder).Error
	return
}
