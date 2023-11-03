package utils

import (
	"encoding/json"
	"net/http"
)

func BuildFailedResponse(w http.ResponseWriter, err error, status int) {
	response := map[string]interface{}{}
	response["status"] = false
	response["msg"] = "falied to process"
	response["data"] = nil
	response["error"] = err.Error()
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func BuildSuccessResponse(w http.ResponseWriter, msg string, data interface{}, status int) {
	response := map[string]interface{}{}
	response["status"] = true
	response["msg"] = msg
	response["data"] = data
	response["error"] = nil
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}
