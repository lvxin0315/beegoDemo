package background

/* 监控第三方平台状态定时执行任务 */
import (
	"github.com/robfig/cron"
	"github.com/astaxie/beego/logs"
)

func Run() {
	c := cron.New()
	//c.AddFunc("*/30 * * * * ?", updateThirdPartyPlatformStatus)
	c.AddFunc("*/10 * * * * ?", testMongodb)
	c.Start()
}

//testMongodb
func testMongodb()  {
	logs.Info("start test background run")
}
