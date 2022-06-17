/**
 * @Author: yangyi
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午2:42
 */

package controller

import (
	"errors"
	"fmt"
	"gin-gorm/dao"
	"gin-gorm/model"
	"gin-gorm/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

//获取用户信息
func UserList(c *gin.Context) {
	//接收参数
	userId := c.Query("id")
	fmt.Println("=====userId====", userId)
	userid, _ := strconv.Atoi(userId)
	fmt.Println("=====userid====", userid)
	data, err := dao.GetUserList(userid)
	if err != nil {
		fmt.Printf("查询错误 %s", err)
		return
	}
	render(c, data, err)
}

//新增用户信息
func Create(c *gin.Context) {
	userId := c.Query("userId")
	userid, _ := strconv.Atoi(userId)
	gender := c.Query("gender")
	genders, _ := strconv.Atoi(gender)
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	bool := dao.Create(userid, genders, username, password, email)
	render(c, bool, nil)
}

//修改用户数据
func UpdateUser(c *gin.Context) {
	id := c.Query("id")
	Id, _ := strconv.Atoi(id)
	username := c.Query("username")
	isUpdate := dao.UpdateUser(Id, username)
	render(c, isUpdate, nil)

}

//删除批量用户
func DelUser(c *gin.Context) {
	id := c.Query("id")
	userId, _ := strconv.Atoi(id)
	isDel := dao.DelUser(userId)
	render(c, isDel, nil)
}

// LoginHandler 登录业务
// @Summary      登录业务
// @Description  登录业务
// @Tags         用户业务接口
// @Accept       application/json
// @Produce      application/json
// @Param        Authorization  header  string           false  "B1earer 用户令牌"
// @Param        object         query   model.LoginForm  false  "查询参数"
// @Security     ApiKeyAuth
//@Success 200 {object} _ResponsePostList
// @Router       /login [POST]
func Login(c *gin.Context) {
	//1.获取请求参数
	var u *model.LoginForm
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err 是不是validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationError类型错误 直接返回
			ResponseError(c, CodeInvalidParams) //请求参数错误
			return
		}
		//validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理-登陆
	user, err := service.Login(u)

	if err != nil {
		zap.L().Error("service.Login failed", zap.String("username", u.UserName), zap.Error(err))
		if errors.Is(err, dao.ErrorUserNotExit) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeInvalidParams)
		return
	}
	//3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id": fmt.Sprintf("%d", user.UserID), //js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})

}
// SignUpHandler 注册业务
// @Summary      注册业务
// @Description  注册业务
// @Tags         用户业务接口
// @Accept       application/json
// @Produce      application/json
// @Param        Authorization  header  string              false  "Bearer 用户令牌"
// @Param        object         query   model.RegisterForm  false  "查询参数"
// @Security     ApiKeyAuth
//@Success 200 {object} _ResponsePostList
// @Router       /signup [POST]
func Signup(c *gin.Context)  {
	//1.获取请求参数
	  var r *model.RegisterForm
	  if err := c.ShouldBindJSON(&r); err !=nil{
         //请求参数有误，直接返回
	  	zap.L().Error("signup with invalid param",zap.Error(err))
	  	fmt.Println("===f===",r)
	  	//判断err 是不是validator.ValidationErrors类型的errors
	  	errs,ok := err.(validator.ValidationErrors)
	  	if !ok {
	  		//非validator.ValidationError类型错误 直接返回
	  		ResponseError(c,CodeInvalidParams)//请求参数错误
			return
		}
		//validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		  return
	  }
	  //业务处理 注册

	  if err := service.SignUp(r); err != nil{
	  	zap.L().Error("service.signup failed",zap.Error(err))
	  	if errors.Is(err,dao.ErrorUserExit){
	  		ResponseError(c,CodeUserExist)
			return
		}
		ResponseError(c,CodeServerBusy)
		  return
	  	if err != nil{
	  		zap.L().Error("register() is failed",zap.Error(err))
	  		ResponseError(c,CodeServerBusy)
			return
		}
	  }

	//返回响应
	ResponseSuccess(c, nil)



}
