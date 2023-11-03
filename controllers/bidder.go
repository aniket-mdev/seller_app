package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sellerapp/database"
	"sellerapp/models"
	"sellerapp/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Bidders struct {
	db *gorm.DB
}

func NewBidders() *Bidders {
	db := database.InitDb()
	db.AutoMigrate(&models.Bidders{})
	return &Bidders{db: db}
}

func (bidder *Bidders) CreateBidder(w http.ResponseWriter, r *http.Request) {
	// read a request Body
	request_body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// extarct a request body into model
	var bidderDTO models.CreateBidderDTO

	if err := json.Unmarshal(request_body, &bidderDTO); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	// validate a model for required fields
	st := validator.New()

	if err := st.Struct(bidderDTO); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	// check bidding is closed or not

	// convert dto to model
	bidder_model := models.Bidders{
		Name:      bidderDTO.Name,
		BidAmount: bidderDTO.BidAmount,
		AdSpaceId: bidderDTO.AdSpaceId,
	}

	// create a bidder for the ad space
	if err := models.CreateBidde(bidder.db, &bidder_model); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "Bidder has been created", bidder_model, http.StatusCreated)
}

func (bidder *Bidders) GetBidders(w http.ResponseWriter, r *http.Request) {
	var bidders []models.Bidders

	err := models.GetBidders(bidder.db, &bidders)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.BuildSuccessResponse(w, "Fetch success", bidders, http.StatusOK)
}

// fetch a bidders by ad space id
func (bidder *Bidders) GetBidderByAdSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ad_space_id"]

	space_id, err := strconv.Atoi(id)

	if err != nil {
		utils.BuildFailedResponse(w, errors.New("invalid param found"), http.StatusBadRequest)
		return
	}

	var bidders []models.Bidders
	if err := models.GetBiddersByAdSpace(bidder.db, &bidders, space_id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BuildFailedResponse(w, err, http.StatusNotFound)
			return
		}
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "Fetch Success", bidders, http.StatusOK)
}

// fetch a all bidding info by bidder name
func (bidder *Bidders) GetBidderByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if name == "" {
		utils.BuildFailedResponse(w, errors.New("provide a required params"), http.StatusBadRequest)
		return
	}

	var bidders []models.Bidders

	if err := models.GetBidder(bidder.db, &bidders, name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BuildFailedResponse(w, err, http.StatusNotFound)
			return
		}
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "Fetch Success", bidders, http.StatusOK)
}

// update a bidde by bidde id
func (bidder *Bidders) UpdateBidder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// converting a string to int
	bidde_id, err := strconv.Atoi(id)
	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	// check bidde exists or not and bidde status as well
	var bidders models.Bidders
	if err := models.GetBiddeById(bidder.db, &bidders, bidde_id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BuildFailedResponse(w, err, http.StatusNotFound)
			return
		}

		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// check bidde status , it should be a open for update bidde
	if bidders.Status != "open" {
		utils.BuildFailedResponse(w, errors.New("can not update a close bidde"), http.StatusBadRequest)
		return
	}

	// read a request body for update and only bidding Amount will be update
	request_body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	// extract request body into DTO
	var update_bidder models.UpdateBiddeDTO
	if err := json.Unmarshal(request_body, &update_bidder); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// set a new bid amount to the bidder model
	bidders.BidAmount = update_bidder.BidAmount

	// send a model to update
	if err := models.UpdateBidder(bidder.db, &bidders); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "update sucesss", bidders, http.StatusOK)
}
