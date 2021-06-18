package main

import (
	"MyMachinesStack/app"
	"MyMachinesStack/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o Sistema...")
	app.InitServer()
	routes.MyRoutes()

}
