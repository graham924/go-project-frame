package main

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/server"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	cmd := server.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
