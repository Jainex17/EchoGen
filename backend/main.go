package main

import "echogen/backend/app"

func main() {
	app.LoadEnv()
	app.AuthInit()
	app.Run()
}
