package global

// pkg/setting为读取configs中的配置信息，此global/setting 声明全局setting变量，供应用程序内部调用

import "github.com/consult98/simple-blog-go/pkg/setting"

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
