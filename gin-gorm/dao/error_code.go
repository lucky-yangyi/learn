/**
 * @Author: yangyi
 * @Description: //定义业务状态码
 * @File:  error_code
 * @Version: 1.0.0
 * @Date: 2022/6/14 下午4:27
 */

package dao

import "errors"

var (
    ErrorUserExit      = errors.New("用户已存在")
    ErrorUserNotExit   = errors.New("用户不已存在")
    ErrorPasswordWrong = errors.New("密码错误")
    ErrorGenIDFailed   = errors.New("创建用户ID失败")
    ErrorInvalidID     = errors.New("无效的ID")
    ErrorQueryFailed   = errors.New("查询数据失败")
    ErrorInsertFailed  = errors.New("插入数据失败")
)
