package main

import (
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Controller"
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Login"
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Signup"
	"net/http"
)

func LoginHandlerGCF(w http.ResponseWriter, r *http.Request) {
	Login.LoginHandler(w, r)
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	Login.LoginHandler(w, r)
}
func main() {
	http.HandleFunc("/", EntryPoint)
	Controller.Auth()
	// Menghubungkan rute HTTP dari package login
	// Mendaftarkan rute HTTP dari package login
	// Mendaftarkan rute HTTP dari package signup
	http.HandleFunc("/Signup", Signup.SignupHandler)
	Login.RegisterLoginRoutes()
	//Mendaftarkan Fungsi GCF
	// Melayani form login

	// Mulai server HTTP
	http.ListenAndServe(":8989", nil)
}
