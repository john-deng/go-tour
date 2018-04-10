package primary

import (
	"fmt"
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

func createdMap(t *testing.T)  {

	dict := make(map[string]int)
	dict["张三"] =3
	dict["cs"] = 4
	age := dict["cs"]
	//当存在的时候exists true  不存在  为false
	age , exists :=dict["cd"]
	assert.Equal(t,map[string]int{"张三":3,"cs":4},dict)
	//删除key
	delete(dict,"cs")
	fmt.Printf("%v\n",dict)
	fmt.Printf("%v\n",age)
	fmt.Printf("%v\n",exists)
	assert.Equal(t,map[string]int{"张三":3},dict)
	assert.Equal(t,false,exists)
	assert.Equal(t,0,age)
}



func forMap(t *testing.T)  {

	dict := make(map[string]int)
	dict["张三"] =3
	dict["cs"] = 4
	for _,v:=range dict{
		fmt.Println(v)
	}
	var names []string
	for key,value :=range dict{
		fmt.Printf("key:%v\n , value:%v\n",key,value)
		names = append(names,key)
	}
	fmt.Println("names:",names)
	assert.Equal(t,[]string{"张三","cs"},names)
	sort.Strings(names)
	assert.Equal(t,[]string{"cs","张三"},names)
}


