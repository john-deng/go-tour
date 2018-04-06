package primary

import (
	"fmt"
	"sort"
)

func createdMap()  {

	dict := make(map[string]int)
	dict["张三"] =3
	dict["cs"] = 4
	age := dict["cs"]
	//当存在的时候exists true  不存在  为false
	age , exists :=dict["cd"]
	//删除key
	delete(dict,"cs")
	fmt.Printf("%v\n",dict)
	fmt.Printf("%v\n",age)
	fmt.Printf("%v\n",exists)
}



func forMap()  {

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
	sort.Strings(names)

}


