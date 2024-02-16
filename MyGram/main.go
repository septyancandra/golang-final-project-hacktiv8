package main

import (
	"Mygram/database"
	"Mygram/routers"
)

func main() {
	config.StartDB()
	r := routers.StartApp()
	r.Run(":8085")
}
