/**
 * @Author: yangyi
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午3:28
 */

package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-gorm/model"
	"gin-gorm/utils"
)

func Login(user *model.User) (err error) {
	//var u *model.User
	originPassword := user.Password
	err = model.GetDB().Raw(`select user_id,username,password from user where username = ?`, user.UserName).Scan(&user).Error
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("查询数据库出错!", err)
		return
	}
	if err == sql.ErrNoRows {
		//用户不存在
		return ErrorUserNotExit
	}
	//生成加密密码与查询到的密码比较
	password := utils.EncryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	return

}

/**
 * @Author yangyi
 * @Description //TODO 检测用户信息是否存在
 * @Date 上午10:47 2022/6/15
 * @Param
 * @return
**/

func CheckUserExist(username string)(err error)  {
	var count int
	if err := model.GetDB().Raw(`select count(user_id) from user where username = ?`,username).Find(&count).Error; err !=nil{
		return err
	}
	if count >0 {
		errors.New("用户已存在")
	}
	return
}

/**
 * @Author yangyi
 * @Description //TODO  注册业务-向数据库中插入一条新数据
 * @Date 上午11:12 2022/6/15
 * @Param
 * @return
**/
func InsertUser(user model.User)(err error)  {
	//1.对密码进行加密
	user.Password = utils.EncryptPassword([]byte(user.Password))
	err = model.GetDB().Exec(`insert into user(user_id,username,password)values(?,?,?)`,user.UserID,user.UserName,user.Password).Error
	return err


}

//获取用户信息
func GetUserList(userId int) (userList []model.User, err error) {
	model.GetDB().Debug().Preload("Orders").Raw(`select id,user_id,username,password,email,gender,create_time,update_time from user where id = ?`, userId).Scan(&userList)
	return
}

//新增用户信息
func Create(userId, gender int, username, password, email string) (err error) {
	model.GetDB().Exec(`insert into user(user_id,gender,username,password,email)values(?,?,?,?,?)`, userId, gender, username, password, email)
	return
}

//修改用户信息
func UpdateUser(Id int, username string) bool {
	model.GetDB().Exec(`update user set username = ? where id = ? `, username, Id)
	return true

}

//删除用户
func DelUser(userId int) bool {
	model.GetDB().Exec(`delete from user where id = ? `, userId)
	return true

}
