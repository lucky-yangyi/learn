/**
 * @Author: yangyi
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午2:52
 */

package model

//用户信息
//type User struct {
//    gorm.Model
//    Username string `json:"username"`  //用户名
//    Password string `json:"password"` //密码
//    Email    string `json:"email"` //邮箱
//    Gender   int `json:"gender"` //年级
//    CreateTime time.Time `json:"createTime"`
//    UpdateTime time.Time `json:"updateTime"`
//    Orders     []*Order
//}

type User struct {
    UserID   uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
    UserName string `json:"username" db:"username"`
    Password string `json:"password" db:"password"`
    AccessToken    string
    RefreshToken   string
}

/**
 * @Author yangyi
 * @Description //TODO  登陆请求参数
 * @Date 上午11:56 2022/6/14
**/
type LoginForm struct {
    UserName string `json:"username" binding:"required" comment:"用户名"`
    PassWord string `json:"password" binding:"required"`
}

/**
 * @Author yangyi
 * @Description //TODO  注册请求参数
 * @Date 上午11:56 2022/6/14
**/
type RegisterForm struct {
    UserName        string `json:"username" binding:"required"`
    Password        string `json:"password" binding:"required"`
    ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
