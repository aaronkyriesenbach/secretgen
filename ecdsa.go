package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetECDSAKeyPair(c *gin.Context) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bytes, _ := x509.MarshalECPrivateKey(privateKey)
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: bytes,
	})

	c.JSON(http.StatusOK, encoded)
}
