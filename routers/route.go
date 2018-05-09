package routers

import (
	"github.com/gorilla/mux"
	"github.com/writetosalman/go-rest-api-boilerplate/http/controllers"
	"github.com/writetosalman/go-rest-api-boilerplate/http/middlewares"
)

func InitialiseRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", 				controllers.Homepage)
	router.HandleFunc("/ping", 			controllers.Ping)
	router.HandleFunc("/login", 			controllers.Login)
	router.HandleFunc("/dashboard", 		middlewares.Authentication(controllers.Dashboard))
	
	return router
}
