package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-gin-video/docs"
	"go-gin-video/initialization"
	"go-gin-video/middleware"
	"net/http"
)

func InitRouter() error {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// 使用Zap的log,自定义Recovery
	engine.Use(initialization.GinLogger, initialization.GinRecovery(false))
	// 解决跨域问题
	engine.Use(middleware.CORS())
	collectRouter(engine)

	// 注册swagger路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := engine.Run(fmt.Sprintf("%s:%s",
		viper.GetString("server.host"),
		viper.GetString("server.port")),
	)
	if err != nil {
		return err
	}
	return nil
}

func collectRouter(engine *gin.Engine) {
	// 静态资源
	static := engine.Group("/static")
	static.StaticFS("/images", http.Dir("./static/resources/images"))
	static.StaticFS("/video", http.Dir("./static/resources/video"))
	// 用户相关路由
	UserRouter(engine)
	CommentRouter(engine)
	CaptchaRouter(engine)
	UploadRouter(engine)
	VideoRouter(engine)
	LikeRouterInit(engine)
	CollectRouterInit(engine)
	FollowRouterInit(engine)
}
