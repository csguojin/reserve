package dal

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/csguojin/reserve/config"
	"github.com/csguojin/reserve/model"
)

func GetDB() (*gorm.DB, *redis.Client) {
	onceDB.Do(createDB)
	onceRDB.Do(createRDB)
	return DB, RDB
}

var (
	DB     *gorm.DB
	onceDB sync.Once

	RDB     *redis.Client
	onceRDB sync.Once
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

	err = db.AutoMigrate(&model.Resv{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Admin{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func createRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("db.redis.host") + ":" + viper.GetString("db.redis.port"),
		DB:   0,
	})

	RDB = rdb
}

const (
	redisTTL = time.Hour
)
