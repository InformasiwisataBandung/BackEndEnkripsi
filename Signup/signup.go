package Signup

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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
			// Cek apakah username sudah ada di database
			if usernameExists(username) {
				http.Error(w, "Username sudah digunakan, silakan coba username lain", http.StatusConflict)
				return
			}

			// Hash the password using bcrypt
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Printf("Gagal mengenkripsi kata sandi: %v", err)
				http.Error(w, "Gagal menyimpan data ke MongoDB", http.StatusInternalServerError)
				return
			}

			user := User{Username: username, Password: string(hashedPassword)}
			collection := client.Database("InformasiWisataBandung").Collection("Users")
			_, err = collection.InsertOne(context.Background(), user)
			if err != nil {
				log.Printf("Gagal menyimpan data ke MongoDB: %v", err)
				http.Error(w, "Gagal menyimpan data ke MongoDB", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/template/login.html", http.StatusSeeOther)
			return
		}
	} else if r.Method == http.MethodGet {
		// Tampilkan formulir pendaftaran untuk metode GET
		http.ServeFile(w, r, "templates/signup.html")
	} else {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
	}

	http.ServeFile(w, r, "templates/signup.html")
}

// Function to check if username exists in the database
func usernameExists(username string) bool {
	collection := client.Database("InformasiWisataBandung").Collection("Users")
	filter := bson.M{"username": username}

	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}
