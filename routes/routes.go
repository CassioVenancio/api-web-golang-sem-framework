package routes

import (
	"net/http"
	"project/m/v2/controllers"
)

// Rotas de acesso
func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
