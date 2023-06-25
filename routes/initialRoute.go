package routes

//path: routes/initialRoute.go

import (
	"training/controllers"
	"training/repository"

	"github.com/gorilla/mux"
)

func InitialRoute() *mux.Router {
	router := mux.NewRouter()

	// Connect to database
	db := repository.NewPostgresRepo()

	// Initialize controller
	userController := controllers.NewUserController(db)

	// auth endpoints
	router.HandleFunc("/login", userController.Login).Methods("POST")
	router.HandleFunc("/register", userController.Register).Methods("POST")

	// Set middleware
	router.Use(loggingMiddleware)
	router.Use(rateLimitMiddleware)

	return router
}
