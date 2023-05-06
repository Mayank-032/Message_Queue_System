package main

import (
	"go-message_queue_system/bootstrap"
	"go-message_queue_system/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Printf("Error: %v, unable to init DB", err)
		return
	}

	r := gin.Default()
	apiGroup := r.Group("/api")
	bootstrap.Init(apiGroup)

	r.Run(":8000")
}