package main

import (
	"echogen/backend/app"
	"echogen/backend/config"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()
	app.AuthInit()
	app.Run()
}
