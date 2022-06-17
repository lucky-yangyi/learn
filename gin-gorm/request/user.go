/**
 * @Author: yangyi
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/6/14 上午11:52
 */

package request


/**
 * @Author yangyi
 * @Description //TODO  登陆请求参数
 * @Date 上午11:56 2022/6/14
**/

type LoginForm struct {
    UserName string `json:"username" binding:"required"`
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
