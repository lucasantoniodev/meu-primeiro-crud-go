package main

import (
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/configuration/env"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/controller/routes"
)

func main() {
	env.LoadEnvSettings()

	routes.RunServer(":8080")
}
