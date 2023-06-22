package setting

import "time"

type ServerSettingS struct { //声明配置属性的结构体，与 config.yaml 配置文件中的参数相同
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct { //在global/setting中实例化结构体变量，以读取初始配置参数
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// 结构体方法，由Setting结构体变量直接调用
func (s *Setting) ReadSection(k string, v interface{}) error { //读取区段配置的配置方法，在main.go中调用，将配置读取到全局变量v中
	err := s.vp.UnmarshalKey(k, v) // k为configs中某个区段名，将其区段中配置内容映射到应用配置的结构体中。
	if err != nil {
		return err
	}

	return nil
}
