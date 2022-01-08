package main

import (
	"billulloSRC/Routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	route := mux.NewRouter()

	Routes.SetBillulloRutas(route)

	serverBillullo := http.Server{
		Addr: ":80",
		Handler: route,
	}

	log.Println(serverBillullo.ListenAndServe())
}
