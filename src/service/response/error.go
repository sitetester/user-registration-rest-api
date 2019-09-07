package response

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorParams struct {
	Status  string
	Message string
}

type ApiResponse struct {
}

func (apiResponse ApiResponse) Error(w http.ResponseWriter, req *http.Request, message string) {
	var params ErrorParams
	params.Status = "Error"
	params.Message = message

	msg, _ := json.Marshal(params)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(msg))
}

func (apiResponse ApiResponse) Success(w http.ResponseWriter, req *http.Request, message string) {
	var params ErrorParams
	params.Status = "Error"
	params.Message = message

	msg, _ := json.Marshal(params)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(msg))
}
