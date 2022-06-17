/**
 * @Author: yangyi
 * @Description:
 * @File:  md5
 * @Version: 1.0.0
 * @Date: 2022/6/14 下午4:47
 */

package utils

import (
    "crypto/md5"
    "encoding/hex"
)

const secret = "huchao.vip"

func EncryptPassword(data []byte) (result string) {
    h := md5.New()
    h.Write([]byte(secret))
    return hex.EncodeToString(h.Sum(data))
}
