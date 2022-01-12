package main

import (
	"billulloSRC/Routes"
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/mailgun/mailgun-go/v4"
	_ "github.com/mailgun/mailgun-go/v4/events"
	"log"
	"net/http"
	"time"
)

func main() {

	r := mux.NewRouter()
	Routes.ConfigRouter(r)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8081",
		Handler:      r,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("tubillullo_com.crt", "HSSL-61dda12abc2af.key"))
}

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Freddy Campos <fcampos@tubillullo.com>",
		"tubillullo backend test",
		"probando funciones en backend!",
		"marthianfred@gmail.com",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
