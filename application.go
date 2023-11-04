package main

import (
	"fmt"
	"net/http"
	"sellerapp/controllers"

	"github.com/gorilla/mux"
)

func main() {
	setupHttpRouter()
}

func setupHttpRouter() {

	adspace := controllers.NewAdSpaceRepo()
	router := mux.NewRouter()

	router.HandleFunc("/adspace", adspace.CreateNewAdSpace).Methods(http.MethodPost)
	router.HandleFunc("/adspace", adspace.GetAdSpaces).Methods(http.MethodGet)
	router.HandleFunc("/adspace/{id}", adspace.GetAdSpace).Methods(http.MethodGet)
	router.HandleFunc("/adspace/{id}", adspace.UpdateAdSpace).Methods(http.MethodPut)
	router.HandleFunc("/adspace/{id}", adspace.DeleteAdSpace).Methods(http.MethodDelete)
	router.HandleFunc("/check_adspace_endtime", adspace.CheckAdSpaceEndTime).Methods(http.MethodGet)

	http.Handle("/", router)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
