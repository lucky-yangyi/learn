/**
 * @Author: yangyi
 * @Description:
 * @File:  resp
 * @Version: 1.0.0
 * @Date: 2022/5/30 下午2:44
 */

package model

type Response struct {
    Code    int         `json:"code"`
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
}
