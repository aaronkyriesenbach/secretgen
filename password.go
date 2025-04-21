package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"net/http"
)

type PasswordRequest struct {
	Names       []string `form:"names[]"`
	Length      int      `form:"length"`
	NumDigits   int      `form:"numDigits"`
	NumSymbols  int      `form:"numSymbols"`
	NoUpper     bool     `form:"noUpper"`
	AllowRepeat bool     `form:"allowRepeat"`
}

func GetPassword(c *gin.Context) {
	request := PasswordRequest{
		Names:       []string{"secretKey"},
		Length:      64,
		NumDigits:   8,
		NumSymbols:  8,
		NoUpper:     false,
		AllowRepeat: true,
	}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	generated := map[string]string{}
	for _, name := range request.Names {
		res, err := password.Generate(request.Length, request.NumDigits, request.NumSymbols, request.NoUpper, request.AllowRepeat)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		generated[name] = res
	}

	c.JSON(http.StatusOK, generated)
}
