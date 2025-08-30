package main

import (
	"echogen/backend/app"
	"echogen/backend/config"
)

func main() {
	config.LoadEnv()
	app.AuthInit()
	app.Run()
}
