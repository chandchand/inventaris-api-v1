package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBUSERNAME string
	DBPASSWD   string
	DBHOST     string
	DBPORT     string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Load .env file")
	}

	config := &Config{
		DBUSERNAME: os.Getenv("DB_USERNAME"),
		DBPASSWD: os.Getenv("DB_PASSWORD"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
	}

	return config
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/inventaris", c.DBUSERNAME, c.DBPASSWD, c.DBHOST, c.DBPORT)
}

func (c *Config) InitDB() (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(c.GetConnectionString()), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}