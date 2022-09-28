package model

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Log() {
	logPath := "tmp"
	os.MkdirAll(logPath, os.ModePerm)

	accessLogFile, _ := os.Create(fmt.Sprintf("%s/access.log", logPath))
	gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

	errorLogFile, _ := os.Create(fmt.Sprintf("%s/error.log", logPath))
	gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stderr)
}
