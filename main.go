package main

import (
	"fmt"
	"os"
	"strings"

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

	router := gin.Default()

	router.GET("/api/*path", func(c *gin.Context) {
		path := strings.TrimPrefix(c.Param("path")+"*", "/")

		entries, err := client.Glob(path)
		if err != nil {
			fmt.Printf("Could not list pattern %q, err: %v\n",
				path, err)
			c.Status(500)
			return
		}

		filenames := []string{}
		for _, entry := range entries {
			filenames = append(filenames, string(entry.Name))
		}

		c.JSON(200, gin.H{
			"files": filenames,
		})
	})

	router.Run()
}
