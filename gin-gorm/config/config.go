/**
 * @Author: yangyi
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/5/30 上午10:39
 */

package config

import (
    "encoding/json"
    "io/ioutil"
	"os"
)

//配置文件结构体
type config struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Ports  string  `json:"ports"`
}

var Config  *config
//打开配置文件
func init() {
	file, err := os.Open("./conf/app.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	by, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
    con :=new(config)
    if err := json.Unmarshal(by,con); err !=nil{

    }
    Config = con
}
