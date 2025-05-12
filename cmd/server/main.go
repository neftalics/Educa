package main

import (
	"server/internal/config"
	"server/internal/content"
)

func main() {
	config.LoadEnv()
	content.GetClientes()
}
