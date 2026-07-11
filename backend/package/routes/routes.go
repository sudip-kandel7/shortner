package routes

import (
	"github.com/gorilla/mux"
)

var RoutesHandler = func(r *mux.Router) {
	r.HandleFunc("/", controllers.SignIn).Methods("POST")
}
