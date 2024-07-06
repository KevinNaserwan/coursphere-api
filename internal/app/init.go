package app

import (
	"github.com/kevinnaserwan/coursphere-api/internal/util/env"
)

type app struct {
	config *env.Config
}

//	@contact.name	Coursphere API
//	@contact.email	kevinnaserwan@gmail.com

//	@securityDefinitions.apikey	UserAuthorization
//	@in							header
//	@name						Authorization
//	@description				User 	Jwt Token Authorization

//	@securityDefinitions.apikey	AdminAuthorization
//	@in							header
//	@name						Authorization
//	@description				Admin	Jwt Token Authorization

// @externalDocs.description	OpenAPI
//
// @externalDocs.url			https://swagger.io/resources/open-api/
func NewApp(config *env.Config) *app {
	return &app{
		config: config,
	}
}
