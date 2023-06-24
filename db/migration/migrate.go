package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sword/api-backend-challenge/db/mysql"
	"github.com/sword/api-backend-challenge/log"
	"sync"
)

var (
	driver database.Driver
	once   sync.Once
)

func NewMigrate() *migrate.Migrate {

	logger := log.NewEntry()

	once.Do(func() {
		db := mysql.GetConn()

		var err error
		driver, err = migratemysql.WithInstance(db, &migratemysql.Config{})
		if err != nil {
			logger.WithError(err).Fatal()
		}

	})

	migration, err := migrate.NewWithDatabaseInstance(
		"file://./db/migration",
		"mysql", driver)
	if err != nil {
		logger.WithError(err).Fatal()
	}

	return migration
}
