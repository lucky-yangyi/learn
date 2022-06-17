/**
 * @Author: yangyi
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2022/6/15 下午2:55
 */

package middleware

import (
    "fmt"
    "gin-gorm/controller"
    "gin-gorm/middleware/jwt"
    "github.com/gin-gonic/gin"
    "strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
        if authHeader == ""{
            controller.ResponseErrorWithMsg(c,controller.CodeInvalidToken,"缺少请求头Auth Token")
            c.Abort()
            return
        }
        //按空格分割
        parts := strings.SplitN(authHeader,"",2)
        if !(len(parts) == 2 && parts[0] == "Bearer"){
            controller.ResponseErrorWithMsg(c,controller.CodeInvalidToken,"Token格式不对")
            c.Abort()
            return
        }
        //parts[1] 是获取到的tokenString,我们使用之前定义好的解析JWT函数来解析它
        mc, err := jwt.ParseToken(parts[1])
        if err != nil{
            fmt.Println(err)
            controller.ResponseError(c,controller.CodeInvalidToken)
            c.Abort()
            return
        }
        //将当前请求的userID信息保存到请求的上下文c中
        c.Set(controller.ContextUserIDKey,mc.UserID)
        c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息

    }

}
