package gormdb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"klever.io/interview/utils"
)

func getConnection(dataBaseName string) *gorm.DB {
	var db *gorm.DB

	switch strings.ToLower(dataBaseName) {
	case "gorm_postgres":
		db = getPostgresConnection()
	default:
		db = getPostgresConnection()
	}

	runMigrations(db)

	return db
}

func getPostgresConnection() *gorm.DB {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gormdb", "function": "getPostgresConnection"})

	port, err := strconv.Atoi(utils.GetEnv("DB_PORT", "5432"))
	if err != nil {
		port = 5432
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASS", "postgres"),
		utils.GetEnv("DB_DATABASE", "upvote_db"),
		port,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logCtx.Fatalf("could not create postgres database connection %v", err)
	}

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	var errorList []error
	logCtx := logrus.WithFields(logrus.Fields{"component": "gormdb", "function": "runMigrations"})
	err := db.AutoMigrate(&Coin{})
	if err != nil {
		errorList = append(errorList, err)
	}

	if len(errorList) > 0 {
		logCtx.Fatalf("could not run migration on database: %v", errorList)
	}
}
