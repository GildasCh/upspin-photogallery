package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gildasch/upspin-photogallery/collection"
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

	router.Static("/static", "./static")

	router.LoadHTMLFiles("templates/index.html")
	router.GET("/s/*path", func(c *gin.Context) {
		filenames, err := fileserver.List(c.Param("path"))
		if err != nil {
			fmt.Println("fileserver.List:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"collection": collection.New(filenames),
		})
	})

	router.GET("/f/*path", func(c *gin.Context) {
		reader, err := fileserver.Get(c.Param("path"))
		if err != nil {
			fmt.Println("fileserver.Get:", err)
			c.Status(http.StatusBadRequest)
			return
		}

		c.Stream(func(w io.Writer) bool {
			_, err := io.CopyN(w, reader, 1024*1024)
			return err == nil
		})
	})

	router.Run()
}
