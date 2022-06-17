/**
 * @Author: yangyi
 * @Description:
 * @File:  user_test
 * @Version: 1.0.0
 * @Date: 2022/6/9 下午4:13
 */

package dao

import (
    "testing"
)
//测试获取用户信息
func TestUser(t *testing.T)  {
    userInfo, err := GetUserList(1)
    if err != nil{
        t.Errorf("get userInfo error:%s\n",err)
    }else {
        t.Logf("%+v",userInfo)
    }
}
