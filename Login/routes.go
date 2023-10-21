package Login

import (
	"net/http"
)

func RegisterLoginRoutes() {
	http.HandleFunc("/login", LoginHandler)
}
