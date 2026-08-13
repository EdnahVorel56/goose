package dialect

import (
	"fmt"
	"github.com/pressly/goose/v3" // Assuming standard import path
)

func GetVersionQuery() string {
	return fmt.Sprintf("SELECT version_id, is_applied FROM %s ORDER BY id DESC", goose.GetQualifiedTableName())
}

func CreateTableQuery() string {
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id SERIAL PRIMARY KEY, version_id BIGINT NOT NULL, is_applied BOOLEAN NOT NULL, tstamp TIMESTAMP DEFAULT now())", goose.GetQualifiedTableName())
}