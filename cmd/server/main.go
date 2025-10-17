package main

import config "github.com/NakaokaGabriel/go/products-projects/configs"

func main() {
	config, _ := config.LoadConfig(".")

	println(config.DBDriver)
}
