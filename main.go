package main

import (
	"github.com/Edilberto-Vazquez/personal-API/libs"
	"github.com/Edilberto-Vazquez/personal-API/routes"
)

func main() {
	libs.Conexion()
	routes.Run()
}
