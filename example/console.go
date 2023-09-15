package main

import (
	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/database"
	"github.com/fwidjaya20/go-framework/facades"
	"github.com/golang-module/carbon/v2"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&database.ServiceProvider{},
	})

	facades.Config().Add("database", map[string]any{
		"connections": map[string]any{
			"postgresql": map[string]any{
				"driver":   "postgresql",
				"host":     facades.Config().Get("DB_HOST", "127.0.0.1"),
				"port":     facades.Config().Get("DB_PORT", 5432),
				"database": facades.Config().Get("DB_DATABASE", "forge"),
				"username": facades.Config().Get("DB_USERNAME", ""),
				"password": facades.Config().Get("DB_PASSWORD", ""),
			},
		},
		"default":  facades.Config().Get("DB_CONNECTION", "postgresql"),
		"dir":      "./out/database",
		"timezone": carbon.UTC,
	})

	facades.App().Boot()
}
