package database

import (
	"fmt"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	sslMode := os.Getenv("SSL_MODE")
	sslRootCert := os.Getenv("SSL_ROOTCERT")
	options := os.Getenv("OPTIONS")

	urlValues := url.Values{"sslmode": []string{sslMode}}
	if sslRootCert != "" {
		urlValues["sslrootcert"] = []string{sslRootCert}
	}
	if options != "" {
		urlValues["options"] = []string{options}
	}
	dsn := url.URL{
		User:     url.UserPassword(dbUsername, dbPassword),
		Scheme:   "postgresql",
		Host:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:     dbTable,
		RawQuery: urlValues.Encode(),
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn.String(), // data source name
		PreferSimpleProtocol: true,         // disables implicit prepared statement usage.
	}), &gorm.Config{})
	if err != nil {
		return db, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	if sqlDB.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
