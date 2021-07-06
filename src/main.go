package main

import (
	"log"
	"net/http"
	

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
)

// LoadEnv /*
func LoadEnv(){

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
}


/*
 Entry point
*/
func main() {

	//Starting the API server
	router := routes.LogRoutes()

	//Load the env file
	LoadEnv()
	http.Handle("/api/", router)
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:4200"}))(router)))

	
	

	

	// log.Println("Listening...")
	 //log.Fatal(http.ListenAndServe(":3000", nil))


}
