/**
 * @Author: yangyi
 * @Description:
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2022/5/31 上午11:16
 */

package model

import "gorm.io/gorm"

//订单结构体
type Order struct {
    gorm.Model
    UserId int `json:"user_id"` //关联用户id
    Price   float64 `json:"price"` //价格
}
