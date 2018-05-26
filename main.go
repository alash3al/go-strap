package main

import (
	"os"
	"strconv"
	"strings"

	"core.globals"
	"core.models"
	"core.routes"
	"github.com/go-redis/redis"
)

func main() {
	// the version number
	globals.LogSuccess("⇨ Starting ... \n")

	// load the global .env file
	globals.Boot()

	// connect to the configured db server
	if strings.ToLower(globals.ENV["DB_ENABLE"]) == "true" {
		models.Boot()
	}

	// connect to redis
	if strings.ToLower(globals.ENV["REDIS_ENABLE"]) == "true" {
		redidb, _ := strconv.Atoi(globals.ENV["REDIS_DB"])
		globals.RedisConn = redis.NewClient(&redis.Options{
			Addr:     globals.ENV["REDIS_HOST"],
			Password: globals.ENV["REDIS_PASSWORD"],
			DB:       redidb,
		})
		if _, err := globals.RedisConn.Ping().Result(); err != nil {
			globals.LogErr("⇨ [Redis] %s \n", err.Error())
			os.Exit(1)
		}
		globals.LogSuccess("⇨ Connected to the redis server ... \n")
	}

	// initalizing the routes server
	routes.Serve()
}
