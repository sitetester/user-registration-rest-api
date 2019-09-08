package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"user-registration-rest-api/src/routes/users"
)

func ManageRoutes() {
	var u users.RouteUsers

	indexHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/users/login", u.LoginHandler).Methods("POST")
	r.HandleFunc("/users/dashboard", u.DashboardHandler).Methods("POST")
	r.HandleFunc("/users/register", u.RegisterHandler).Methods("POST")

	r.HandleFunc("/users/uploadImage", u.UploadImageHandler).Methods("POST")
	http.Handle("/", r)

	http.ListenAndServe(":8090", nil)
	fmt.Printf("Starting server for testing HTTP POST...\n")
}
