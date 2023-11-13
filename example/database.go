package main

import (
	"fmt"

	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/database"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

func readAllDatabaseTables(session *gorm.DB) []string {
	// Raw SQL query to get the list of tables in the current database
	rows, err := session.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Rows()
	if err != nil {
		panic("Failed to execute the query")
	}
	defer rows.Close()

	// Scan the results into a slice of strings
	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic("Failed to scan row")
		}
		tables = append(tables, tableName)
	}

	return tables
}

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

	// to run the apps
	app := facades.App()
	app.Boot()

	// do database thingy
	session := app.GetDatabase()

	tables := readAllDatabaseTables(session)
	fmt.Println(tables)
}
