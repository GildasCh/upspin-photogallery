package main

import (
	"fmt"
	"os"

	"github.com/gildasch/upspin-photogallery/files"
	"github.com/gin-gonic/gin"
	"upspin.io/client"
	"upspin.io/config"
	_ "upspin.io/transports"
)

func main() {
	cfg, err := config.FromFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	client := client.New(cfg)
	fileserver := &files.Server{Accesser: client}

	router := gin.Default()

	router.GET("/api/*path", func(c *gin.Context) {
		filenames, err := fileserver.List(c.Param("path"))
		if err != nil {
			fmt.Println(err)
			c.Status(500)
			return
		}

		c.JSON(200, gin.H{
			"files": filenames,
		})
	})

	router.Run()
}
