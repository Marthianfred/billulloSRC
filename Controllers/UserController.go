package Controllers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){

}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("hola estoy es probando el beta..."))
}

func CreateUser(w http.ResponseWriter, r *http.Request){

}

func DeleteUser(w http.ResponseWriter, r *http.Request){

}

func UpdateUser(w http.ResponseWriter, r *http.Request){

}
