package models

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"core.globals"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Boot connect to the database
func Boot() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@%s/%s?%s",
		globals.ENV["DB_USERNAME"],
		globals.ENV["DB_PASSWORD"],
		globals.ENV["DB_HOST"],
		globals.ENV["DB_DATABASE"],
		globals.ENV["DB_OPTIONS"],
	)

	if strings.ToLower(globals.ENV["DB_NEED_SCHEME"]) == "true" {
		dsn = fmt.Sprintf("%s://%s", globals.ENV["DB_DRIVER"], dsn)
	}

	globals.LogWarn("⇨ [DB] %s \n", dsn)

	globals.DBConn, err = sql.Open(globals.ENV["DB_DRIVER"], dsn)

	if err != nil {
		globals.LogErr("⇨ [DB] %s \n", err.Error())
		os.Exit(1)
	}

	if err = globals.DBConn.Ping(); err != nil {
		globals.LogErr("⇨ [DB] %s \n", err.Error())
		os.Exit(1)
	}

	globals.LogSuccess("⇨ Connected to the database server ... \n")

	migrate()

	globals.LogSuccess("⇨ Migrated the database table(s) successfully ... \n")
}

func migrate() {

}
