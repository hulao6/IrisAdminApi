package models

import (
	"IrisYouQiKangApi/config"
	"IrisYouQiKangApi/redis"
	"IrisYouQiKangApi/tools"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB     *gorm.DB       //mysql
	Tools  *tools.Tools   //tools
	Config *config.Config //config
	//Modb   *mongodb.Mongodb //mongodb
	Redis *redis.Redis //redis
)

func init() {
	Config = config.New()
	Tools = tools.New()
	//Modb = mongodb.New(Config.Mongodb.Connect)
	Redis = redis.New(Config.Redis.Connect, Config.Redis.DB, Config.Redis.MaxIdle, Config.Redis.MaxActive)

	var (
		err error
	)
	DB, err = gorm.Open("mysql", Config.Mysql.Connect)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}

	//DB.AutoMigrate(&OauthToken{}, &Users{})

}

/**
 * 用于api输出信息
 */
type ApiJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}
