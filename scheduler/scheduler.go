package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

func setupScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(1).Minute().Do(func() {
		fmt.Println("Scheduler has been run...")
		checkAdSpaceEndTime()
	})

	scheduler.StartBlocking()
}

type AdSpace struct {
	ID          int         `json:"ID"`
	CreatedAt   time.Time   `json:"CreatedAt"`
	UpdatedAt   time.Time   `json:"UpdatedAt"`
	DeletedAt   interface{} `json:"DeletedAt"`
	Description string      `json:"description"`
	BasePrice   int         `json:"base_price"`
	Status      string      `json:"status"`
	EndTime     string      `json:"end_time"`
}

type AdSpacesResponse struct {
	Data   []AdSpace   `json:"data"`
	Error  interface{} `json:"error"`
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
}

func checkAdSpaceEndTime() error {
	var url = "http://localhost:8080/check_adspace_endtime"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			return
		}
	}()

	if response.StatusCode == http.StatusOK {
		var adSpaces AdSpacesResponse

		respoonse_body, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		if err := json.Unmarshal(respoonse_body, &adSpaces); err != nil {
			return err
		}

		for i := range adSpaces.Data {
			err = close_adspace_status(adSpaces.Data[i])
			if err != nil {
				continue
			}
		}

	}

	return errors.New("some thing wents wrong")
}

func close_adspace_status(adspace AdSpace) error {
	ad_id := strconv.Itoa(adspace.ID)

	var url = "http://localhost:8080/adspace/" + ad_id

	adspace.Status = "close"
	requestBody, err := json.Marshal(&adspace)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(requestBody))

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			return
		}
	}()

	if response.StatusCode != http.StatusOK {
		return errors.New("someting wents wrong")

	}
	return nil
}

func main() {
	setupScheduler()
}
