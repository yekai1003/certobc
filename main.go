package main

import (
	"github.com/yekai1003/certobc/routes"

	_ "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var EchoObj *echo.Echo //echo框架对象全局定义

func staticFile() {

	EchoObj.Static("/", "static")       //根目录设置
	EchoObj.Static("/static", "static") //全路径处理

}

func main() {
	EchoObj = echo.New()
	EchoObj.Use(middleware.Logger()) //安装日志中间件
	EchoObj.Use(middleware.Recover())

	EchoObj.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	staticFile()

	EchoObj.GET("/ping", routes.PingHandler)
	EchoObj.POST("/register", routes.Register)
	EchoObj.POST("/login", routes.Login)
	EchoObj.POST("/issue", routes.Issue)
	EchoObj.POST("/query", routes.Query)
	EchoObj.POST("/course", routes.AddCourse)
	EchoObj.Logger.Fatal(EchoObj.Start(":8080")) //启动服务
}
