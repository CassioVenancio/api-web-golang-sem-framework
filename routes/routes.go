package routes

import (
	"net/http"
	"project/m/v2/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index) // Rota de acesso
}
