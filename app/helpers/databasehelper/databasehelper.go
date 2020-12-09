package databasehelper

import (
	"errors"
	"fmt"
	"mycart/app/helpers/confighelper"
	"mycart/app/helpers/filehelper"
	"mycart/app/helpers/loghelper"
	"mycart/app/models"

	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBTomlConfig -
type DBTomlConfig struct {
	DBHosts map[string]DBConfiguration
}

// DBConfiguration -
type DBConfiguration struct {
	HostName   string
	Server     string
	Port       int
	Username   string
	Password   string
	Database   string
	IsDefault  bool
	IsDisabled bool
}

var once sync.Once

var dbConnections map[string]*gorm.DB
var defaultHost string

func init() {
	dbConnections = make(map[string]*gorm.DB)
}

// InitDatabases creates database connection and stores in the pool of dbConnections
func InitDatabases() error {

	var dbErr error
	once.Do(func() {
		dbConfigs := DBTomlConfig{}

		_, err := confighelper.InitConfig(models.MYSQLConfigPath, &dbConfigs)
		if err != nil {
			loghelper.LogError("error while seting database configurations:", err)
			dbErr = err
		}

		f, err := filehelper.FilePointer(models.DBLogPath)
		if err != nil {
			fmt.Println("Can not open dblog file ::", err)
			dbErr = err
		}
		newLogger := logger.New(
			log.New(f, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // Disable color
			},
		)
		for hostName, config := range dbConfigs.DBHosts {

			dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Server, config.Port, config.Database)

			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: newLogger,
			})
			if err != nil {
				loghelper.LogError("error while connecting database:", err)
				dbErr = err
			}

			dbConnections[hostName] = db

			dbConnections[hostName].AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
			// dbConnections[hostName].AutoMigrate(&models.Room{})
			// dbConnections[hostName].AutoMigrate(&models.RatePlan{})

			if config.IsDefault {
				defaultHost = hostName
			}
		}
	})

	return dbErr
}

// GetConnectionByHostName returns the database object of provided hostName
func GetConnectionByHostName(hostName string) (*gorm.DB, error) {

	if hostName == "" {
		connection, found := dbConnections[defaultHost]
		if !found {
			loghelper.LogError("Connection not found!")
			return nil, errors.New("Connection not found")
		}
		return connection, nil
	}
	connection, found := dbConnections[hostName]
	if !found {
		loghelper.LogError("Connection not found for ", hostName)
		return nil, errors.New("Connection not found for " + hostName)
	}
	return connection, nil

}
