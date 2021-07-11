package server

import (
	"net/http"
	"os"
	"os/signal"
	"log"

	"github.com/gin-gonic/gin"
)

// config struct
type Config struct {
	Port string
	Debug bool
}

// server struct
type Server struct {
	*gin.Engine
}

// perform healthcheck
func healthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

// new server
func New() Server {
	// binding.Validator = new(validator.DefaultValidator)
	server := gin.Default()
	server.GET("/", healthCheck)
	return Server{server}
}


// start server
func Start(e *Server, cfg *Config) {

	s := &http.Server{
		Addr:    cfg.Port,
		Handler: e.Engine,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.Close(); err != nil {
			log.Println("Failed To ShutDown Server", err)
		}
		log.Println("Shuting Down Server")
	}()

	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server Closed After Interruption")
		} else {
			log.Println("Unexpected Server Shutdown. err:", err)
		}
	}
}
