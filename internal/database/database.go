package database

import (
	"fmt"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbSSLRootCert := os.Getenv("DB_SSLROOTCERT")
	dbOptions := os.Getenv("DB_OPTIONS")

	urlValues := url.Values{"sslmode": []string{dbSSLMode}}
	if dbSSLRootCert != "" {
		urlValues["sslrootcert"] = []string{dbSSLRootCert}
	}
	if dbOptions != "" {
		urlValues["options"] = []string{dbOptions}
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
