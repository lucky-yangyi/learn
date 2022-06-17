/**
 * @Author: yangyi
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/5/30 上午10:01
 */

package main

import (
    "gin-gorm/config"
    "gin-gorm/router"
)

//@title bluebell_backend
//@version 1.0
//@description gin-gorm测试
//@termsOfService http://swagger.io/terms/
//
//@contact.name author：@yangyi
//@contact.url http://www.swagger.io/support
//@contact.email support@swagger.io
//
//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
//@host 127.0.0.1:8081
//@BasePath /api/v1/

func main()  {
    //go dao.GetDog()
   engine := router.Router()
   engine.Run(config.Config.Ports)

}
