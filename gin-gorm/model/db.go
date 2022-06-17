/**
 * @Author: yangyi
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午2:03
 */

package model

import (
    "fmt"
    "gin-gorm/config"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

var db  *gorm.DB

func init()  {
    //dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.Config.UserName,config.Config.PassWord,config.Config.Host + config.Config.Port,config.Config.Name)), &gorm.Config{
        NamingStrategy:                           schema.NamingStrategy{
            SingularTable: true,
        },
    })
    if err != nil{
       fmt.Println(err)
        return
    }
    //自动迁移
    GetDB().AutoMigrate(&GirlGOD{},&Dog{},&Info{})
    //info1 := Info{
    //    Age:   1,
    //    Money: 1,
    //}
    //info2 := Info{
    //    Age:   2,
    //    Money: 2,
    //}
    //dog1 := Dog{
    //    Name: "dog1",
    //    Info: info1,
    //}
    //dog2 := Dog{
    //    Name: "dog2",
    //    Info: info2,
    //}
    //girl1 := GirlGOD{
    //    Name: "girl1",
    //    Dog:  []Dog{dog1},
    //}
    //girl2 := GirlGOD{
    //    Name: "girl2",
    //    Dog:  []Dog{dog2},
    //}
    //db.Debug().Create(&girl1)
    //db.Debug().Create(&girl2)


}
func GetDB() *gorm.DB {
    if db == nil {
        return nil
    }
    return db
}
