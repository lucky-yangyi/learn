/**
 * @Author: yangyi
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午2:40
 */

package router

import (
    "gin-gorm/controller"
    "github.com/gin-gonic/gin"
   _ "gin-gorm/docs"
    //ginSwagger "github.com/swaggo/gin-swagger"
    //"github.com/swaggo/gin-swagger/swaggerFiles"

    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
    r := gin.New()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/api/v1/login",controller.Login)
    r.POST("/api/v1/signup",controller.Signup)
    //r.GET("/api/v1/userList",controller.UserList)
    //r.POST( "/api/v1/Create",controller.Create)
    //r.POST( "/api/v1/updateUser",controller.UpdateUser)
    //r.DELETE("/api/v1/delUser",controller.DelUser)
    return r

}
