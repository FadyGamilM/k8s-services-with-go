package transport

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	return gin.Default()
}

func CreateServer(r *gin.Engine, PORT string) *http.Server {
	return &http.Server{
		Addr:    PORT,
		Handler: r,
	}
}

func InitServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed initiating the server")
		log.Println(err)
		os.Exit(1)
	} else {

		log.Println("🥇 Server is up and running ..")
	}
}
