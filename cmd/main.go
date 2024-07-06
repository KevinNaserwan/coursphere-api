package main

import (
	"os"

	App "github.com/kevinnaserwan/coursphere-api/internal/app"
	"github.com/kevinnaserwan/coursphere-api/internal/util/env"
)

func main() {
	config := env.LoadConfig()

	os.Setenv("TZ", "Asia/Jakarta")
	app := App.NewApp(config)

	app.StartServer()
}
