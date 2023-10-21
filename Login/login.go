package Login

import (
	"context"
	"github.com/whatsauth/watoken"
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

var Privatekey = "56e4eb16f428e82cea21e5bceed2b078c0955ce1b8509631369dab20e1a952180a9ea5fae87b3895fba98c2b138c336ccfba886b0823fd774415ccc9394ae159"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Retrieve hashed password from MongoDB based on the username
	hashedPassword, err := getHashedPassword(username)
	if err != nil {
		http.Error(w, "Gagal mencari kredensial", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		http.Error(w, "Login gagal", http.StatusUnauthorized)
		return
	}

	// If login is successful, generate a PASETO token
	tokenstring, _ := watoken.Encode(username, Privatekey)
	w.Write([]byte(tokenstring))
}
func getHashedPassword(username string) (string, error) {
	// Koneksi ke MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://MigrasiData:Salman123456.@cluster0.ot8qmry.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Koneksikan ke koleksi "Users" di database "InformasiWisataBandung"
	collection := client.Database("InformasiWisataBandung").Collection("Users")

	// Lakukan query untuk mendapatkan dokumen dengan username yang sesuai
	var result User
	filter := bson.M{"username": username}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		// Handle error, misalnya jika dokumen tidak ditemukan
		log.Printf("Gagal mencari kredensial: %v", err)
		return "", err
	}

	// Kembalikan kata sandi terenkripsi dari dokumen yang sesuai
	return result.Password, nil
}
