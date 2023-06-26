/*
Copyright © 2023 NAME HERE ivorfn@gmail.com
*/
package main

import (
	"github.com/sword/api-backend-challenge/api"
)

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	api.NewServer().Run()
}
