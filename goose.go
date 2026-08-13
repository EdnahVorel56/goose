package goose

import (
	"strings"
	"fmt"
)

var tableName = "goose_db_version"

func SetTableName(name string) {
	tableName = name
}

func TableName() string {
	return tableName
}

// GetQualifiedTableName returns the schema-qualified table name for SQL queries.
func GetQualifiedTableName() string {
	parts := strings.Split(tableName, ".")
	quoted := make([]string, len(parts))
	for i, part := range parts {
		quoted[i] = fmt.Sprintf("\"%s\"", strings.ReplaceAll(part, "\"", "\"\""))
	}
	return strings.Join(quoted, ".")
}