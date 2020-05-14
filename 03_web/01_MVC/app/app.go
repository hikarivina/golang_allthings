package app

import (
	"allthinggo/03_web/01_MVC/controllers"
	"net/http"
)

func StartApp() {
	http.HandlerFunc("/user", controllers.GetUser())
}
