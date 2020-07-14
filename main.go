package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appModule, err := appModuleInjector()

	if err != nil {
		fmt.Println(err)
	}

	if err := appModule.Bootstrap(); err != nil {
		panic(err)
	}
}
