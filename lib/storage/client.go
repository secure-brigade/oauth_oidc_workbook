package storage

import (
	"context"
	"fmt"
	"oauth-az/ent"
	"oauth-az/lib/env"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func NewMySQL() *ent.Client {
	conf := ConfigDB("MID_TERM_MANAGEMENT")
	drv, err := sql.Open(dialect.MySQL, conf.DSN())
	//client, err := ent.Open(dialect.MySQL, conf.DSN())

	if err != nil {
		panic(fmt.Sprintf("failed connecting to storage: %v", err))
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	client := ent.NewClient(ent.Driver(drv))

	if env.IsDebug() {
		client = client.Debug()
	}

	return client
}

func NewSQLite() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(fmt.Sprintf("failed connecting to storage: %v", err))
	}

	// localなのでmigrationも走らせちゃう
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	return client
}
