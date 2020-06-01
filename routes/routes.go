package routes

import (
	"net/http"

	"github.com/joaocvr/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
