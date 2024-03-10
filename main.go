package main

import (
	"log"

	"github.com/Drafteame/taskrun/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal("Error: ", err)
	}
}
