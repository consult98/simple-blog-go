package main

import (
	"log"
	"net/http"
	"time"

	"github.com/consult98/simple-blog-go/global"
	"github.com/consult98/simple-blog-go/internal/model"
	"github.com/consult98/simple-blog-go/internal/routers"
	"github.com/consult98/simple-blog-go/pkg/setting"
	"github.com/gin-gonic/gin"
)

// 项目配置参数初始化方法init()，在main之前自动执行
// 一个项目启动的顺序：全局变量初始化 => init方法 => main方法
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 读取项目启动参数配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	// 根据app初始化参数创建log文件
	err = setting.CreateIfNotExists(global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt, 0755)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting) //注意此处不能用:=  因为需要将数据库引擎存入全局变量
	if err != nil {
		return err
	}

	return nil
}
