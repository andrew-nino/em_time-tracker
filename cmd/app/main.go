package main

import "github.com/andrew-nino/em_time-tracker/internal/app"

//	@title			Effective Mobile API
//	@version		1.0.0
//	@description	API Server for test work

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	app.Run()
}
