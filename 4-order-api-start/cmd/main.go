package main

import (
	"order-api/config"
	"order-api/dbl"
)

func main() {
	cfg := config.LoadConfig()
	_ = dbl.NewDB(cfg)
}
