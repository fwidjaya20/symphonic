package main

import (
	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/database"
	"github.com/fwidjaya20/go-framework/foundation"
	"github.com/golang-module/carbon/v2"
)

func main() {
	foundation.App.GetConfig().Add("app.providers", []ContractFoundation.ServiceProvider{
		&database.ServiceProvider{},
	})

	foundation.App.GetConfig().Add("database", map[string]any{
		"connections": map[string]any{
			"postgresql": map[string]any{
				"driver":   "postgresql",
				"host":     foundation.App.GetConfig().Get("DB_HOST", "127.0.0.1"),
				"port":     foundation.App.GetConfig().Get("DB_PORT", 5432),
				"database": foundation.App.GetConfig().Get("DB_DATABASE", "forge"),
				"username": foundation.App.GetConfig().Get("DB_USERNAME", ""),
				"password": foundation.App.GetConfig().Get("DB_PASSWORD", ""),
			},
		},
		"default":  foundation.App.GetConfig().Get("DB_CONNECTION", "postgresql"),
		"dir":      "./out/database",
		"timezone": carbon.UTC,
	})

	foundation.App.Boot()
}
