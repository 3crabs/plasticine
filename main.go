package main

import (
	"plasticine/db"
	"plasticine/server"
)

func main() {
	server.NewServer(":8080", db.NewDB()).Run()
}
