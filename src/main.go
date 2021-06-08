package main

import (
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
	"github.com/joho/godotenv"
	
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
	http.Handle("/api/", router)


	//Load the env file
	LoadEnv()

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))


}
