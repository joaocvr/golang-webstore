package routes

import (
	"net/http"

	"github.com/joaocvr/golang-webstore/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
