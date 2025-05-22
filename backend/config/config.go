package config

import (
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type Config struct {
	SuperTokens supertokens.TypeInput
	Database    DatabaseConfig
	Server      ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Port string
}

func newSuperTokensConfig() supertokens.TypeInput {
	return supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "https://try.supertokens.com",
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "SuperTokens Demo App",
			APIDomain:     "http://localhost:3001",
			WebsiteDomain: "http://localhost:3000",
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
			dashboard.Init(nil),
		},
	}
}

// DefaultConfig returns a default configuration for the application
func DefaultConfig() Config {
	return Config{
		SuperTokens: newSuperTokensConfig(),
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "postgres",
			Database: "cryptodiary",
		},
		Server: ServerConfig{
			Port: "3001",
		},
	}
}
