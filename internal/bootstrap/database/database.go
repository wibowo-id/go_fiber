package database

import (
	"github.com/rs/zerolog"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// setup database with gorm
type Database struct {
	DB  *gorm.DB
	Log zerolog.Logger
	Cfg *config.Config
}

type Seeder interface {
	Seed(*gorm.DB) error
	Count() (int, error)
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

// connect database
func (_db *Database) ConnectDatabase() {
	conn, err := gorm.Open(mysql.Open(_db.Cfg.DB.Mysql.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database !")
	} else {
		_db.Log.Info().Msg("Connected the database successfully !")
	}

	_db.DB = conn
}

// shutdown database
func (_db *Database) ShutdownDatabase() {
	sqlDB, err := _db.DB.DB()
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database !")
	} else {
		_db.Log.Info().Msg("Shutdown the database successfully !")
	}
	sqlDB.Close()
}

// migrate models
func (_db *Database) MigrateModels() {
	if err := _db.DB.AutoMigrate(
		Models()...,
	); err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to migrate the database!")
	}
}

// list of models for migration
func Models() []interface{} {
	return []interface{}{
		schema.ForgotPassword{},
		schema.Permission{},
		schema.Role{},
		schema.RolePermission{},
		schema.User{},
		schema.Version{},
		schema.Wilayah{},
	}
}

// seed data
func (_db *Database) SeedModels(seeder ...Seeder) {
	for _, seed := range seeder {
		count, err := seed.Count()
		if err != nil {
			_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
		}

		if count == 0 {
			if err := seed.Seed(_db.DB); err != nil {
				_db.Log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
			}

			_db.Log.Info().Msg("Seeded the database succesfully!")
		} else {
			_db.Log.Info().Msg("Database is already seeded!")
		}
	}

	_db.Log.Info().Msg("Seeded the database succesfully!")
}
