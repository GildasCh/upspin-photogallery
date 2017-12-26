package main

import (
	"fmt"
	"os"

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
	entries, err := client.Glob("augie@upspin.io/Images/Augie/***")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range entries {
		fmt.Println(entry.Name)
	}
}
