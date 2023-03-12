package initialization

import (
	"go-gin-video/service/mongoDBService"
	"go-gin-video/service/redisService"
	"go.uber.org/zap"
	"log"
)

func InitAll() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	if err := InitViper(false); err != nil {
		log.Println("Viper初始化失败: " + err.Error())
		return
	}

	if err := InitLogger(); err != nil {
		log.Println("Zap初始化失败: " + err.Error())
		return
	}

	if err := InitMysql(); err != nil {
		zap.L().Error("Mysql连接失败: " + err.Error())
		return
	}

	if err := InitGorm(); err != nil {
		zap.L().Error("Gorm初始化失败: " + err.Error())
		return
	}

	if err := redisService.InitRedis(); err != nil {
		zap.L().Error("Redis连接失败: " + err.Error())
		return
	}

	if err := mongoDBService.InitMongoDb(); err != nil {
		zap.L().Error("MongoDB连接失败: " + err.Error())
	}
	if err := InitCasbin(); err != nil {
		zap.L().Error("casbin初始化错误: " + err.Error())
	}
}
