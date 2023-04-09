package main

import (
	"backend/cmd"
	"backend/config"
)

// @title Stand with Dorayaki API
// @version 2.0
// @description Stand with Dorayaki API Documentation

// @contact.name Kadek Surya Mahardika
// @contact.email kadeksuryam@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
func main() {
	config.Init()
	cmd.Init()
}
