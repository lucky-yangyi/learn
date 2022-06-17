/**
 * @Author: yangyi
 * @Description: //TODO 存放登陆业务逻辑代码
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/6/14 下午3:39
 */

package service

import (
    "fmt"
    "gin-gorm/dao"
    "gin-gorm/middleware/jwt"
    "gin-gorm/model"
)

func Login(p *model.LoginForm) (user *model.User,err error)  {
    user = &model.User{
        UserName: p.UserName,
        Password: p.PassWord,
    }
    if err := dao.Login(user); err != nil {
        return nil,err
    }
    // 生成JWT
    //return jwt.GenToken(user.UserID,user.UserName)
    fmt.Println("===user===",user.UserID)
    atoken,rtoken, err := jwt.GenToken(user.UserID,user.UserName)
    if err != nil {
        return
    }
    user.AccessToken = atoken
    user.RefreshToken = rtoken
    return
}

/**
 * @Author yangyi
 * @Description //TODO  //存放注册的逻辑
 * @Date 上午10:21 2022/6/15
 * @Param
 * @return
**/

func SignUp(p * model.RegisterForm)(err error)  {
    // 1.判断用户存不存在
    if err := dao.CheckUserExist(p.UserName); err !=nil{
        //查询数据库错误
        return  err
    }
    //2。构造一个User实例
    u := model.User{
        UserID:       1,
        UserName:     p.UserName,
        Password:     p.Password,
    }
    //3.保存数据库
    return  dao.InsertUser(u)



}


