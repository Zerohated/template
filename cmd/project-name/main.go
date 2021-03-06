package main

import (
	"syscall"

	conf "github.com/Zerohated/project-name/configs"
	"github.com/Zerohated/project-name/internal/controller"
	"github.com/Zerohated/project-name/pkg/dao"
	"github.com/Zerohated/tools/pkg/logger"

	"github.com/fvbock/endless"
	"github.com/robfig/cron/v3"

	"github.com/gin-gonic/gin"
)

var (
	schedule *cron.Cron
	log      = logger.Logger
	config   = conf.Config
)

func init() {
	// Connect DB
	dbConf := config.PostgresConf
	err := dao.ConnectPG(dbConf.Host, dbConf.Port, dbConf.User, dbConf.DBName, dbConf.Password)
	if err != nil {
		log.Warnln(err.Error())
	}
}

func main() {
	if config.Stage == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	basicAuth := gin.BasicAuth(gin.Accounts{config.BasicAuth.Account: config.BasicAuth.Password})
	// generalLimiter := rate.NewLimiter(rate.Limit(config.Limiter.Limit), config.Limiter.Burst)
	// limiterMiddle := middleware.LimiterMiddle(generalLimiter)

	c := controller.NewController()
	// Echo received message
	app.POST("/echo", c.EchoHandler)

	// 定时任务
	if schedule != nil {
		schedule.Stop()
	}
	schedule = cron.New()
	// schedule.AddFunc("0 4 * * *", func() { c.DealTodos() })
	// schedule.AddFunc("0 8 * * *", func() { c.HandleDailyReport() })
	schedule.Start()

	server := endless.NewServer(config.Port, app)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	server.ListenAndServe()
}
