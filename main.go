package main

import (
	_ "github.com/SoMWbRa/rest-api/docs"
	"github.com/SoMWbRa/rest-api/oauth"
	"github.com/SoMWbRa/rest-api/server"
	"log"
)

// @title REST API ECHO

func main() {
	oauth.InitConfig()
	err := server.InitDatabase("data.db")
	if err != nil {
		log.Fatalln(err)
	}
	e := server.InitiateServer()
	e.Logger.Fatal(e.Start(":3000"))
}
