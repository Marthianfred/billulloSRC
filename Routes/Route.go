package Routes

import (
	"billulloSRC/Controllers"
	"github.com/gorilla/mux"
)

func SetBillulloRutas(r *mux.Router)  {

	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/login}", Controllers.Login).Methods("POST")

	subRouter.HandleFunc("/users", Controllers.CreateUser).Methods("POST")
	subRouter.HandleFunc("/users/{id}", Controllers.GetUser).Methods("GET")
	subRouter.HandleFunc("/users/{id}", Controllers.UpdateUser).Methods("PUT")
	subRouter.HandleFunc("/users/{id}", Controllers.DeleteUser).Methods("DELETE")
}