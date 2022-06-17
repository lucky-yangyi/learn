/**
 * @Author: yangyi
 * @Description:
 * @File:  dog
 * @Version: 1.0.0
 * @Date: 2022/5/31 下午1:57
 */

package model

import "gorm.io/gorm"

type Dog struct {
    gorm.Model
    Name      string
    GirlGodID uint
    Info      Info
}
