// Package http handles the http webserver
package http

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// NewHTTPServer returns gin http server
func NewHTTPServer(ginMode string) *gin.Engine {
	if ginMode == "release" {
		gin.SetMode(ginMode)
	} else if ginMode == "test" {
		gin.SetMode(ginMode)
	} else {
		gin.SetMode("debug")
	}

	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Content-Disposition", "X-Filename", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Disposition", "X-Filename"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}
