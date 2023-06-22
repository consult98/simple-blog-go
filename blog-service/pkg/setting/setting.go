package setting //使用 viper 读取 config.yaml 中的默认配置，封装后供应用程序使用

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   //配置文件名称
	vp.AddConfigPath("configs/") //配置文件路径
	vp.SetConfigType("yaml")     //配置文件类型
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
