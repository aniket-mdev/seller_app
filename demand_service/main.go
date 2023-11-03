package main

import (
	"fmt"
	"net/http"
	"sellerapp/controllers"

	"github.com/gorilla/mux"
)

func main() {
	setupHttpRouter()
	fmt.Println("Demand Sevice has been run....")
}

func setupHttpRouter() {

	router := mux.NewRouter()
	bidd_cont := controllers.NewBidders()

	router.HandleFunc("/bidder", bidd_cont.CreateBidder).Methods(http.MethodPost)
	router.HandleFunc("/bidder", bidd_cont.GetBidders).Methods(http.MethodGet)
	router.HandleFunc("/bidder/{ad_space_id}", bidd_cont.GetBidderByAdSpace).Methods(http.MethodGet)

	// bidder name
	router.HandleFunc("/bidder_by_name/{name}", bidd_cont.GetBidderByName).Methods(http.MethodGet)

	// bidder id PK
	router.HandleFunc("/bidder/{id}", bidd_cont.UpdateBidder).Methods(http.MethodPut)

	http.Handle("/", router)

	fmt.Println("Server is listening on :9000")
	http.ListenAndServe(":9000", nil)
}
