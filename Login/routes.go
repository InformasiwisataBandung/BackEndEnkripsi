package Login

import (
	"net/http"
)

func RegisterLoginRoutes() {
	http.HandleFunc("/Login", LoginHandler)
}
