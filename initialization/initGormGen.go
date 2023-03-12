package initialization

import (
	"go-gin-video/service/mysqlService"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDb *gorm.DB

func InitGorm() error {
	var err error
	gormDb, err = gorm.Open(mysql.New(mysql.Config{Conn: mysqlDb}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./service/mysqlService",
		ModelPkgPath: "./domain/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	})
	g.UseDB(gormDb)
	g.ApplyBasic(g.GenerateAllTable()...)
	// 开始生成数据库对应内容
	g.Execute()
	mysqlService.SetDefault(gormDb)
	return nil
}
