package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var QmUser = new(qmUser)

type qmUser struct{}

func (r *qmUser) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("qmUser").Use(middleware.OperationRecord())
		group.POST("createQmUser", apiQmUser.CreateQmUser)
		group.DELETE("deleteQmUser", apiQmUser.DeleteQmUser)
		group.DELETE("deleteQmUserByIds", apiQmUser.DeleteQmUserByIds)
		group.PUT("updateQmUser", apiQmUser.UpdateQmUser)
		group.PUT("adminChangePassword", apiQmUser.AdminChangePassword)
	}
	{
		group := private.Group("qmUser")
		group.GET("findQmUser", apiQmUser.FindQmUser)
		group.GET("getQmUserList", apiQmUser.GetQmUserList)
		group.GET("GetUserInfo", apiQmUser.GetUserInfo)
		group.GET("logout", apiQmUser.Logout)
	}

	{
		group := public.Group("qmUser")
		group.POST("register", apiQmUser.Register)
		group.POST("login", apiQmUser.Login)
		group.GET("getQmUserPublic", apiQmUser.GetQmUserPublic)
	}
}
