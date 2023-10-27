package main

import (
	"godocker/transport"
	"log"
	"os"
)

// func init() {
// 	databaseUrl := os.Getenv("DATABASE_URL")
// 	if databaseUrl == "" {
// 		content, err := ioutil.ReadFile(os.Getenv("DATABASE_URL_FILE"))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		databaseUrl = string(content)
// 	}

// 	errDB := database.InitDB(databaseUrl)
// 	if errDB != nil {
// 		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
// 	} else {
// 		log.Println("DATABASE CONNECTED 🥇")
// 	}

// }

func main() {

	router := transport.CreateRouter()

	PORT := os.Getenv("SERVER_PORT")
	if PORT == "" {
		log.Println("⛔ PORT IS NOT PASSED ⛔")
		PORT = "5000"
	}

	server := transport.CreateServer(router, PORT)
	transport.InitServer(server)
}
