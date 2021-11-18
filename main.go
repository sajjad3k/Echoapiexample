package main

import "github.com/sajjad3k/echoapione/handler"

func main() {
	server := handler.SetRoutes()

	server.Logger.Fatal(server.Start(":8080"))
}
