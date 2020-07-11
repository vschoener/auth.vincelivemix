package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// This file is not mandatory, specially in production so we don't deal with the error
	// I'd prefer using the autoload import but main and env are not in the same path
	godotenv.Load("../.env")

	appModule, err := appModuleInjector()

	if err != nil {
		fmt.Println(err)
	}

	if err := appModule.Bootstrap(); err != nil {
		panic(err)
	}
}
