package Signup

import (
	"context"
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"net/http"
)

type SignupPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GCFPostHandler(w http.ResponseWriter, r *http.Request) {
	// Memastikan request adalah POST
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	// Mendekode JSON payload dari request
	var payload SignupPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Gagal mendekode payload JSON", http.StatusBadRequest)
		return
	}

	// Proses payload sesuai dengan kebutuhan
	// Contoh: Anda dapat menggunakan payload.Username dan payload.Password

	// Kemudian, Anda dapat mengembalikan respons yang sesuai
	// Contoh: Mengembalikan respons JSON
	response := map[string]string{"message": "Pendaftaran berhasil"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func init() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/Signup", GCFPostHandler); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v", err)
	}
}
