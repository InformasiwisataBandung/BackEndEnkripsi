package Signup

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var client *mongo.Client

func init() {
	// Inisialisasi koneksi MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://MigrasiData:Salman123456.@cluster0.ot8qmry.mongodb.net/")
	client, _ = mongo.Connect(context.Background(), clientOptions)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "" && password != "" {
			user := User{Username: username, Password: password}
			collection := client.Database("InformasiWisataBandung").Collection("Users")
			_, err := collection.InsertOne(context.Background(), user)
			if err != nil {
				log.Printf("Gagal menyimpan data ke MongoDB: %v", err)
				http.Error(w, "Gagal menyimpan data ke MongoDB", http.StatusInternalServerError)
				return
			}

			// Setelah sign-up berhasil, redirect ke halaman login atau halaman lain yang sesuai
			http.Redirect(w, r, "/Login", http.StatusSeeOther)
			return
		}
	}

	// Jika metode bukan POST atau data tidak valid, tampilkan halaman sign-up
	http.ServeFile(w, r, "templates/signup.html")
}
