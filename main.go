package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/password", GetPassword)
	router.GET("/ecdsa", GetECDSAKeyPair)
	router.Run()
}
