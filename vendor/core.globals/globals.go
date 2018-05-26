package globals

import (
	"os"

	"database/sql"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	// ENV Environemnt Vars
	ENV map[string]string

	// DBConn Database connection
	DBConn *sql.DB
)

var (
	// LogErr Error Logger
	LogErr = color.New(color.FgRed).PrintfFunc()

	// LogWarn Warning Logger
	LogWarn = color.New(color.FgYellow).PrintfFunc()

	// LogSuccess Success Logger
	LogSuccess = color.New(color.FgGreen).PrintfFunc()

	// Version the version of the system
	Version = "1.0.0"
)

// Boot initialize the environment
func Boot() {
	// our error var
	var err error

	// get the .env filename
	dotenvfilename := ".env"
	if len(os.Args) > 1 {
		dotenvfilename = os.Args[1]
	}

	// load the .env file
	ENV, err = godotenv.Read(dotenvfilename)
	if err != nil {
		LogErr("⇨ [ENV] %s \n", err.Error())
		return
	}

	// Suucceeded!
	LogSuccess("⇨ Loaded the `.env` file  from (%s) ... \n", dotenvfilename)
}
