package primary

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
)

type User struct{
	Name string `name`
	Age int `age`
}

func TestJsonToText(t *testing.T)  {
	var u User
	h:=`{"name":"张三","age":15}`
	err:=json.Unmarshal([]byte(h),&u)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println(u)
		assert.Equal(t, 15 , u.Age)
		assert.Equal(t, "张三" , u.Name)
	}
}


func TestTagToJson(t *testing.T)  {
	var u User
	u.Name = "张三"
	u.Age = 25
	newJson,err:=json.Marshal(&u)
	if err!=nil {

	}else{
		log.Println(string(newJson))
		assert.Equal(t, `{"Name":"张三","Age":25}`, string(newJson))
		Info.Println("info")
		Warning.Println("警告")
		Error.Println("错误")
	}
}


