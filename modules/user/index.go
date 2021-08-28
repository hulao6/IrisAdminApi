package user

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/web"
)

// Party 调试模块
func Party() web.WebModule {
	handler := func(index iris.Party) {
		// index.Use(middleware.JwtHandler().Serve, middleware.New().ServeHTTP, middleware.OperationRecord())
		index.Get("/", GetAllUsers).Name = "用户列表"
		index.Get("/{id:uint}", GetUser).Name = "用户详情"
		index.Post("/", CreateUser).Name = "创建用户"
		index.Post("/{id:uint}", UpdateUser).Name = "编辑用户"
		index.Delete("/{id:uint}", DeleteUser).Name = "删除用户"
		index.Get("/logout", Logout).Name = "退出"
		index.Get("/clear", Clear).Name = "清空 token"
		// index.Get("/expire", controllers.Expire).Name = "刷新 token"
		// index.Get("/profile", controllers.Profile).Name = "个人信息"
		// index.Post("/change_avatar", controllers.ChangeAvatar).Name = "修改头像"
	}
	return web.NewModule("/users", handler)
}
