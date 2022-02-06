package migration

import (
	"context"
	"fmt"
	"log"
	"oauth-az/lib/env"
	"oauth-az/lib/storage"
	"os"
	"time"

	"entgo.io/ent/examples/fs/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	env.SetupVars()
	client := storage.NewMySQL()

	defer client.Close()

	if env.IsDevelop() || env.IsLocal() {
		if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		return
	}

	// Dump migration changes to an SQL script.
	t := time.Now()
	f, err := os.Create(fmt.Sprintf("migrations/%s_migrate.sql", t.Format("20060102150405")))
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()
	if err := client.Schema.WriteTo(ctx, f, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
