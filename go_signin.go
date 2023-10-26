package main

import (
	"github.com/InformasiwisataBandung/BackEndEnkripsi/Login"
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	Login.LoginHandler(w, r)
}
