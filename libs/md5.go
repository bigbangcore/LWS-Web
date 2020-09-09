/**********************************************
** @Des: This file ...
** @Author: zhongjin
** @Date:   2020-07-10 09:58:25
** @Last Modified by:   zhongjin
** @Last Modified time: 2020-07-10 09:58:25
***********************************************/

package libs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var emailSample = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

//Md5 function is used for MD5 encryption
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//SizeFormat is used for size format
func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n++
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

//IsEmail is used to determine whether a string is email address
func IsEmail(b []byte) bool {
	return emailSample.Match(b)
}

//Password is used to generate password
func Password(len int, strPWD string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "lws6688"
	if strPWD != "" {
		defaultPwd = strPWD
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

//MD5 ,生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//GetRandomString ,生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
