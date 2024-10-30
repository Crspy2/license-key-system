package database

import (
	"crspy2/licenses/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Database struct {
	Staff    *Staff
	Sessions *Session
	Users    *User
	Products *Product
	Licenses *License
}

var Client *Database

func newDatabase() *Database {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	i, err := gorm.Open(postgres.Open(config.Conf.Database.URI), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		log.Fatalf("Unable to establish database connection: %#v", err)
	}

	db := &Database{
		Staff:    newStaff(i),
		Sessions: newSessions(i),
		Users:    newUser(i),
		Products: newProduct(i),
		Licenses: newLicense(i),
	}
	return db
}

func (db *Database) createTables() {
	migrateTables(
		db.Staff,
		db.Sessions,
		db.Users,
		db.Products,
		db.Licenses,
	)
}

func migrateTables(tables ...table) {
	for _, t := range tables {
		if err := t.schema(); err != nil {
			panic(err)
		}
	}
}

func ConnectToDatabase() {
	Client = newDatabase()
	Client.createTables()
}
