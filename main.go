package main

import (
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Controller"
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Login"
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Signup"
	"net/http"
)

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/login.html")
		} else {
			http.Error(w, "Metode tidak diizinkancoy", http.StatusMethodNotAllowed)
		}
	})

	// Mulai server HTTP
	http.ListenAndServe(":8989", nil)
}
