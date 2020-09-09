/**********************************************
** @Des: This file ...
** @Author: zhongjin
** @Date:   2020-07-31  09:58:25
** @Last Modified by:   zhongjin
** @Last Modified time: 2020-07-31 09:58:25
***********************************************/

package libs

import (
	"encoding/json"
)

//Acc ,used to test
type Acc struct {
	Name string `json:"user_name"`
	ID   int32  `json:"user_id"`
	Age  uint32 `json:"user_age"`
}

//T , interface
type T interface{}

//Struct2Json ,convert struct to json
func Struct2Json(mo T) string {
	jsonBytes, err := json.Marshal(mo)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

//JSON2Struct is used to convert json to struct
/*
eg:
account := libs.Acc{}
acc := libs.J2Struct(jsonStr, &account)
re, ok := acc.(*libs.Acc)
if ok {
	fmt.Println(*re)
}
*/
func JSON2Struct(str string, t T) bool {
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return false
	}
	return true
}

//JSON2Map ,convert json to map
func JSON2Map(jsonStr string, mapResult map[string]T) bool {
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		return false
	}
	return true
}

//Map2JSON ,convert map to json
func Map2JSON(mapInstances map[string]T) string {
	jsonBytes, err := json.Marshal(mapInstances)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
