package Routes

import (
	"billulloSRC/Controllers"
	"github.com/gorilla/mux"
)

func ConfigRouter(router *mux.Router) {

	router.HandleFunc("/", Controllers.Home)

	apiRoute := router.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/login", Controllers.Login).Methods("POST")
	apiRoute.HandleFunc("/users", Controllers.CreateUser).Methods("POST")
	apiRoute.HandleFunc("/users/{id}", Controllers.GetUser).Methods("GET")
	apiRoute.HandleFunc("/users/{id}", Controllers.UpdateUser).Methods("PUT")
	apiRoute.HandleFunc("/users/{id}", Controllers.DeleteUser).Methods("DELETE")

}
