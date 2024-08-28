package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/api/route"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting video server...")
	r := gin.Default()

	route.Setup(r)
	gin.SetMode(gin.DebugMode)

	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
