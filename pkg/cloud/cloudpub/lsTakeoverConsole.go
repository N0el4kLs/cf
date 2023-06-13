package cloudpub

import (
	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/cloud"
	"github.com/teamssix/cf/pkg/util/database"
)

func LsTakeoverConsole(provider string) {
	TakeoverConsoleCache := database.SelectTakeoverConsoleCache(provider)
	if len(TakeoverConsoleCache) == 0 {
		log.Info("未找到控制台接管信息 (No console takeover information found)")
	} else {
		var (
			header = []string{"云服务商 (Provider)", "用户名 (User Name)", "密码 (Password)", "控制台登录地址 (Login URL)", "Access Key Id", "Access Key Secret", "接管时间 (Takeover Time)"}
			data   [][]string
		)
		for _, v := range TakeoverConsoleCache {
			data = append(data, []string{
				v.Provider,
				v.UserName,
				v.Password,
				v.LoginUrl,
				v.ConsoleAccessKeyId,
				v.ConsoleAccessKeySecret,
				v.CreateTime,
			})
		}
		var td = cloud.TableData{Header: header, Body: data}
		cloud.PrintTable(td, "控制台接管信息 (Console takeover information)")
	}
}
