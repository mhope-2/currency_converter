package server

import (
	"net/http"
	"os"
	"os/signal"
	"log"
	"fmt"
	"io"

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

	// write server logs to file
	logfile, err := os.Create("logs/server.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)

	router := gin.Default()	

	
	router.GET("/", healthCheck)
	return Server{router}
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
