/**
 * @Author: yangyi
 * @Description:
 * @File:  dog_info
 * @Version: 1.0.0
 * @Date: 2022/5/31 下午1:58
 */

package model

import "gorm.io/gorm"

type Info struct {
    gorm.Model
    Age   int
    Money int
    DogID uint
}
