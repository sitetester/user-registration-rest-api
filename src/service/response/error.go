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

type ErrorResponse struct {
}

func (errorResponse ErrorResponse) Error(w http.ResponseWriter, req *http.Request, message string) {
	var params ErrorParams
	params.Status = "Error"
	params.Message = message

	msg, _ := json.Marshal(params)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(msg))
}

func (errorResponse ErrorResponse) Success(w http.ResponseWriter, req *http.Request, message string) {
	var params ErrorParams
	params.Status = "Error"
	params.Message = message

	msg, _ := json.Marshal(params)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(msg))
}
