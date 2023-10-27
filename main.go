package main

import (
	"fmt"
	"godocker/transport"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	PORT = fmt.Sprint("0.0.0.0:", PORT)
	log.Println("The PORT is => ", PORT)
	if PORT == "" {
		log.Println("⛔ PORT IS NOT PASSED ⛔")
		PORT = "0.0.0.0:5000"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "it works !")
	})

	server := transport.CreateServer(router, PORT)
	transport.InitServer(server)
}
