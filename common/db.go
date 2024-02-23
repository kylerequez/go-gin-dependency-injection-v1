package common

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kylerequez/go-gin-dependency-injection-v1/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var host string = os.Getenv("DB_HOST")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var name string = os.Getenv("DB_NAME")
	var port string = os.Getenv("DB_PORT")
	var mode string = os.Getenv("DB_SSLMODE")
	var timezone string = os.Getenv("DB_TIMEZONE")
	var dsnTemplate = os.Getenv("DB_URL")

	if host == "" ||
		user == "" ||
		password == "" ||
		name == "" ||
		port == "" ||
		mode == "" ||
		timezone == "" ||
		dsnTemplate == "" {
		return errors.New("please enter the database variables needed")
	}

	dsn := fmt.Sprintf(dsnTemplate, host, user, password, name, port, mode, timezone)
	log.Println(":::-::: Connecting to " + dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println(":::-::: Successfully connected to DSN")
	DB = db
	return nil
}

func CloseDB() {
	postgres, err := DB.DB()
	if err != nil {
		log.Println(":::-::: Error in closing DB")
	}
	log.Println(":::-::: Closing the DB")
	postgres.Close()
}

func GetDB() *gorm.DB {
	log.Println(":::-::: Retrieving the DB")
	return DB
}

func MigrateDB() error {
	log.Println(":::-::: Migrating the the DB")
	err := DB.AutoMigrate(
		&types.User{},
	)
	if err != nil {
		log.Println(":::-::: Error in migrating the DB")
		return err
	}
	log.Println(":::-::: Successfully migrated the DB")
	return nil
}

func DropAllTables() error {
	log.Println(":::-::: Dropping all tables of the DB")
	err := DB.Migrator().DropTable(
		&types.User{},
	)
	if err != nil {
		log.Println(":::-::: Error in dropping all the tables of the DB")
		return err
	}
	log.Println(":::-::: Dropped all tables of the DB")
	return nil
}
