package dal

import (
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/model"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func createDB() {
	host := viper.GetString("db.mysql.host")
	port := viper.GetString("db.mysql.port")
	user := viper.GetString("db.mysql.username")
	pass := viper.GetString("db.mysql.password")
	database := viper.GetString("db.mysql.database")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect mysql db:" + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Room{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Seat{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	once.Do(createDB)
	return DB
}
