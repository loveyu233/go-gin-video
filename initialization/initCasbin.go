package initialization

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"log"
)

var enforcer *casbin.Enforcer

func InitCasbin() error {
	var err error
	enforcer, err = casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")
	if err != nil {
		zap.L().Error("初始化casbin错误，err：" + err.Error())
		return err
	}
	return nil
}

func Check(sub, obj, act string) bool {
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		log.Println(err)
	}
	return ok
}

func Role(role int) string {
	switch role {
	case 1:
		return "user"
	case 2:
		return "admin"
	case 3:
		return "root"
	default:
		return ""
	}
}
