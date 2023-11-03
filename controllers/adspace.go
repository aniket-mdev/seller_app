package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sellerapp/database"
	"sellerapp/models"
	"sellerapp/utils"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AdSpaceRepo struct {
	Db *gorm.DB
}

func NewAdSpaceRepo() *AdSpaceRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.AdSpace{})
	return &AdSpaceRepo{Db: db}
}

func (repo *AdSpaceRepo) CreateNewAdSpace(w http.ResponseWriter, r *http.Request) {
	var adSpace models.CreateAdSpaceDTO

	request_body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(request_body, &adSpace); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	st := validator.New()
	if err := st.Struct(adSpace); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	new_time, err := time.Parse("2006-01-02 15:04", adSpace.EndTime)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}
	fmt.Println("New Time : ", new_time)
	fmt.Println("New Time Local : ", new_time.Local(), time.Now())

	new_adSpace := models.AdSpace{
		Description: adSpace.Description,
		BasePrice:   adSpace.BasePrice,
		Status:      adSpace.Status,
		EndTime:     adSpace.EndTime,
	}

	err = models.CreateAdSpace(repo.Db, &new_adSpace)
	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "New Ad Space has been created", adSpace, http.StatusAccepted)
}

func (repo *AdSpaceRepo) GetAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpace []models.AdSpace
	err := models.GetAdSpaces(repo.Db, &adSpace)
	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "Fetch success", adSpace, http.StatusOK)
}

func (repository *AdSpaceRepo) GetAdSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	int_id, err := strconv.Atoi(id)

	if err != nil {
		utils.BuildFailedResponse(w, errors.New("invalid param "), http.StatusBadRequest)
		return
	}

	var adSpace models.AdSpace
	err = models.GetAdSpace(repository.Db, &adSpace, int_id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BuildFailedResponse(w, err, http.StatusNotFound)
			return
		}

		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.BuildSuccessResponse(w, "Fetch Success", adSpace, http.StatusOK)
}

func (repository *AdSpaceRepo) UpdateAdSpace(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	int_id, err := strconv.Atoi(id)

	if err != nil {
		utils.BuildFailedResponse(w, errors.New("invalid params found"), http.StatusBadRequest)
		return
	}

	var adSpace models.AdSpace
	err = models.GetAdSpace(repository.Db, &adSpace, int_id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BuildFailedResponse(w, err, http.StatusNotFound)
			return
		}

		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	request_body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// read a body for update
	var update_space_dto models.CreateAdSpaceDTO

	if err := json.Unmarshal(request_body, &update_space_dto); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// validate a request body
	st := validator.New()
	if err := st.Struct(update_space_dto); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusBadRequest)
		return
	}

	// format a end time string to time
	//new_time, err := time.Parse("2006-01-06 15.00", update_space_dto.EndTime)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	// now make a new obj  to update using old one
	adSpace.Description = update_space_dto.Description
	adSpace.EndTime = update_space_dto.EndTime
	adSpace.BasePrice = update_space_dto.BasePrice
	adSpace.Status = update_space_dto.Status

	err = models.UpdateAdSpace(repository.Db, &adSpace)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.BuildSuccessResponse(w, "Update Success", adSpace, http.StatusOK)
}

func (repo *AdSpaceRepo) DeleteAdSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	int_id, err := strconv.Atoi(id)

	if err != nil {
		utils.BuildFailedResponse(w, errors.New("invalid param found"), http.StatusBadRequest)
		return
	}

	var adSpace models.AdSpace
	err = models.DeleteAdSpace(repo.Db, &adSpace, int_id)

	if err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "AdSpace deleted successfully", nil, http.StatusOK)
}

func (repo *AdSpaceRepo) CheckAdSpaceEndTime(w http.ResponseWriter, r *http.Request) {
	var adSpace []models.AdSpace

	if err := models.CheckAdSpaceEndTime(repo.Db, &adSpace); err != nil {
		utils.BuildFailedResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.BuildSuccessResponse(w, "fetch success", adSpace, http.StatusOK)
}
