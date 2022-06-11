package db

import (
	"os"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"

	"github.com/LNMMusic/goapi/config"
)

// Instance
type PsqlInstance struct {
	Db	*gorm.DB
}
var Psql PsqlInstance

// Method
func (p *PsqlInstance) ConnectClient() {
	// URI
	uri := config.EnvGet("PSQLDB_URI")
	if len(uri) == 0 {
		uri = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
							config.EnvGet("PSQLDB_HOST"), config.EnvGet("PSQLDB_PORT"),
							config.EnvGet("PSQLDB_USER"), config.EnvGet("PSQLDB_PSWD"),
							config.EnvGet("PSQLDB_NAME"), config.EnvGet("PSQLDB_SSL"))
	}

	// Connection
	client, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
		os.Exit(2)
	}

	// Instance
	p.Db = client

	log.Printf("Connected!")
}