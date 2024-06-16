package main

import (
	"courseProject/db"
	"courseProject/internal"
)

func main() {
	dbDriver, err := db.InitNewDriver()
	if err != nil {
		panic(err)
	}
	service := internal.NewComplexService(dbDriver)
	handler := internal.NewComplexHandler(service)
	engine := handler.RegisterUrls()

	err = engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
