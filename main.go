package main

import (
	"database/sql"
	"fmt"
	"godocker/transport"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

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

	// Connection parameters
	host := "postgresql" // Service name of the StatefulSet
	port := 5432         // Container port specified in the StatefulSet YAML
	user := "postgres"   // Default user for PostgreSQL
	password := "password"
	dbname := "postgres"

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Println("error trying to ping to the databse => ", err)
		os.Exit(1)
	}

	router := transport.CreateRouter()

	PORT := os.Getenv("SERVER_PORT")
	PORT = fmt.Sprint("0.0.0.0:", PORT)
	log.Println("The PORT is => ", PORT)
	if PORT == "" {
		log.Println("⛔ PORT IS NOT PASSED ⛔")
		// PORT = "0.0.0.0:5000"
		os.Exit(1)
	}

	router.GET("/", func(c *gin.Context) {
		var tm time.Time

		err = db.QueryRow("SELECT NOW() AS now;").Scan(&tm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		c.JSON(http.StatusOK, tm)
	})

	server := transport.CreateServer(router, PORT)
	transport.InitServer(server)
}
