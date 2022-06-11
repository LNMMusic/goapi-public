package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"

	"github.com/LNMMusic/goapi/config"
)

// INSTANCE
type RedisInstance struct {
	Db		*redis.Client
	Ctx 	context.Context
}
var Redis RedisInstance
// var Ctx = context.Background()


// CONNECTION
func (r *RedisInstance) ConnectClient() {
	// URI
	uri := config.EnvGet("REDISDB_URI")
	if len(uri) == 0 {
		uri = fmt.Sprintf("%s:%s", config.EnvGet("REDISDB_HOST"), config.EnvGet("REDISDB_PORT"))
	}

	// Connection
	client := redis.NewClient(&redis.Options{
		Addr:		uri,
		Password:	"",
		DB:			0,
	})
	// Instance
	r.Db = client
	r.Ctx= context.Background()

	// Test
	_, err := r.Db.Ping(r.Ctx).Result()
	if err != nil {
		panic(err)
	}
}