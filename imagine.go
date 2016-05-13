package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.Static("/", "./images")
	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		fileExt := filepath.Ext(header.Filename)

		b := make([]byte, 40)
		rand.Read(b)
		randomName := base64.URLEncoding.EncodeToString(b)
		out, err := os.Create("./images/" + randomName + strings.ToLower(fileExt))

		if err != nil {
			log.Print(err)
			c.Err()
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Print(err)
			c.Err()
		}

		c.String(http.StatusOK, randomName+strings.ToLower(fileExt))
	})

	router.Run(":" + port)
}
