package main

import (
	"MyMachinesStack/app"
	"MyMachinesStack/routes"
	"fmt"
)

func main (){
	fmt.Println("Iniciando o Sistema...\n")

	app.InitServer()
	fmt.Println("Chegou aqui")
	routes.MyRoutes()


}
