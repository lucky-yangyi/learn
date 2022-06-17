/**
 * @Author: yangyi
 * @Description:
 * @File:  dog
 * @Version: 1.0.0
 * @Date: 2022/5/31 下午2:48
 */

package dao

import (
    "fmt"
    "gin-gorm/model"
)

func GetDog()  {
    var girls []model.GirlGOD
    model.GetDB().Debug().Preload("Dog").Preload("Dog.Info").Find(&girls)
    for _,v := range girls{

       for i := range v.Dog{
           fmt.Println(v.Name,v.Dog[i].Name,v.Dog[i].Info.Money)
       }
    }
    //fmt.Println("=====girls====",girls)


}
