package primary

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*基本数据结构*/
/*******************************************************************************************************/
func typeMotify(d string)  string{
	d = "test"
	return d
}

func typeTransmit(tt *testing.T)  {
	s := "tttttt"
	t := typeMotify(s)
	fmt.Printf("s:%v\n",s)
	fmt.Printf("t:%v\n",t)
	assert.Equal(tt,"tttttt",s)
	assert.Equal(tt,"test",t)
}


/*****************************************************************************************************************/
/*map传递值之间
引用类型传递
*/
func typeMapTransmit(t *testing.T)  {
	m := make(map[string]int8)
	m["cs"] = 12
	fmt.Printf("map传递之前的值为: %v\n",m)
	assert.Equal(t,map[string]int8{"cs":12},m)
	typeMapMotify(m)
	fmt.Printf("map传递之后的值为: %v\n",m)
	assert.Equal(t,map[string]int8{"cs":12,"陈爽":25},m)
}
func typeMapMotify(m map[string]int8) {
	m["陈爽"] = 25
}

/*****************************************************************************************************************/

type person struct {
	age int "Age"

	name string "Name"
}

func typePerson()  {
	//var p person
	jim := person{12,"Jim"}
	chenshuang := person{name:"chen",age:12}
	fmt.Printf("jim:%v\n",jim)
	fmt.Printf("chenshuang:%v\n",chenshuang)
}






