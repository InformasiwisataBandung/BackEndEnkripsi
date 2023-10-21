package Login

import (
	"context"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Privatekey = "56e4eb16f428e82cea21e5bceed2b078c0955ce1b8509631369dab20e1a952180a9ea5fae87b3895fba98c2b138c336ccfba886b0823fd774415ccc9394ae159"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if checkCredentials(username, password) {
		// inpo kalo login berhasil mendapatkan token
		tokenString, err := watoken.Encode(username, Privatekey)
		if err != nil {
			http.Error(w, "Gagal menghasilkan token", http.StatusInternalServerError)
			return
		}
		// Kirim token sebagai respons
		w.Write([]byte(tokenString))
	} else {
		// Login gagal
		http.Error(w, "Login gagal", http.StatusUnauthorized)
	}
}

func checkCredentials(username, password string) bool {
	clientOptions := options.Client().ApplyURI("mongodb+srv://MigrasiData:Salman123456.@cluster0.ot8qmry.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("InformasiWisataBandung").Collection("Users")
	var result User
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		// Tangani kesalahan dan kembalikan false jika kredensial tidak ditemukan
		log.Printf("Gagal mencari kredensial: %v", err)
		return false
	}

	// Periksa apakah password cocok dengan yang ada di MongoDB
	if result.Password == password {
		// Kredensial sesuai, maka kembalikan true
		return true
	}

	// Password tidak cocok, kembalikan false
	return false
}
