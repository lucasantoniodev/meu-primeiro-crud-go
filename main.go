package main

import (
	"fmt"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/configuration/env"
	"os"
)

func main() {
	env.LoadEnvSettings()

	fmt.Println(os.Getenv("TEST"))
}
