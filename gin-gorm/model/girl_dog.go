/**
 * @Author: yangyi
 * @Description:
 * @File:  girl_dog
 * @Version: 1.0.0
 * @Date: 2022/5/31 下午1:55
 */

package model

import (
    "gorm.io/gorm"
)

type GirlGOD struct {
    gorm.Model
    Name string
    Dog  []Dog
}


