package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/api/route"
	"github.com/ricardomdesa/videostr/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting video server...")
	env := config.NewEnv()
	r := gin.Default()

	route.Setup(env, r)
	gin.SetMode(gin.DebugMode)
	
	if err := http.ListenAndServe(env.Port, r); err != nil {
		log.Fatal(err)
	}
}
