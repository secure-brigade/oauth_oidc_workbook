package schema

import "entgo.io/ent/dialect"

var textSchema = map[string]string{
	dialect.SQLite:   "text",
	dialect.Postgres: "text",
	dialect.MySQL:    "varchar(384)",
}

var timeSchema = map[string]string{
	dialect.SQLite:   "timestamp",
	dialect.Postgres: "timestamptz",
	dialect.MySQL:    "datetime(3)",
}

var boolSchema = map[string]string{
	dialect.SQLite:   "boolean",
	dialect.Postgres: "boolean",
	dialect.MySQL:    "tinyint(1)",
}
