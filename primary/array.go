package primary

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)
/*********************************************************************************************************************/
func statement(t *testing.T)  {
	var array_list [5] int
	var array_string [5] string
	var array_float [5] float32
	fmt.Printf("array: %v\n", array_list)
	fmt.Printf("array_string: %v\n",array_string)
	fmt.Printf("array_float: %v\n",array_float)
	array_list_eqaul := [5]int {0,0,0,0,0}
	assert.Equal(t,array_list_eqaul,array_list)
	array_string_equal := [5]string {"","","","",""}
	assert.Equal(t,array_string_equal,array_string)
}

/*********************************************************************************************************************/
func initialize(t *testing.T)  {
	var arrayInt [5] int =[5]int{1,1,1,1,1}
	var arrayString [5] string = [5]string{"","","",""}
	var arrayFloat [5] float32= [5]float32{1,1,1,1}
	fmt.Printf("array_int: %v\n",arrayInt)
	fmt.Printf("32小数: %f\n",arrayFloat)
	fmt.Printf("arrayString : %s\n",arrayString)
	assert.Equal(t,[5]float32{1,1,1,1},arrayFloat)
	assert.Equal(t,[5]int{1,1,1,1,1},arrayInt)
	assert.Equal(t,[5]string{"","","",""},arrayString)
}

/*********************************************************************************************************************/
func GoInitialize(t *testing.T)  {
	arrayInt := [5]int{1,2,3,4,5}
	//没有定义长度的赋值
	arrayInt1 :=[...]int{1,2,3}
	fmt.Printf("简单的初始化数组: %v\n",arrayInt)
	fmt.Printf("不需要长度的初始化: %v\n", arrayInt1)
	assert.Equal(t, [5]int{1,2,3,4,5},arrayInt)
	assert.Equal(t,[...]int{1,2,3},arrayInt1)
}

/*********************************************************************************************************************/
func IndexArray(t *testing.T)  {
//索引从0开始 1：1    第二个位置赋值为1
	arrayInt :=[5]int{1:1,3:2}
	arrayString :=[6]string{1:"11" , 3:"Clare"}
	fmt.Printf("%v\n",arrayInt)
	fmt.Printf("索引string类型: %s/n",arrayString)
	assert.Equal(t, [5]int{1:1,3:2}, arrayInt)
	assert.Equal(t,[6]string{1:"11" , 3:"Clare"},arrayString)
}
/*********************************************************************************************************************/
func updateArray(t *testing.T)  {
	fmt.Printf("修改数组中的某个值")
	arrayInt := [...]int{1,2,3,4,5}
	fmt.Printf("当前数组为: %v\n",arrayInt)
	arrayInt[4]= 12
	fmt.Printf("修改后的数组为: %v\n",arrayInt)
	assert.Equal(t,[...]int{1,2,3,12,5},arrayInt)
}


/*********************************************************************************************************************/
func forArray(t *testing.T)  {
	fmt.Printf("循环求得数组索引的值为")
	arrayInt := [...]int{1,2,3,4,5}
	for i:=0 ; i<5 ;i++ {
		fmt.Printf("索引的值为: %d\n , 当前索引的值为: %d\n",i ,arrayInt[i])
	}

	fmt.Printf("循环rang 求得的值")
	for i, v :=range arrayInt{
		fmt.Printf("索引的值为: %d\n , 当前索引的值为: %d\n",i ,v)
	}





}

/*********************************************************************************************************************/

/*数组之间赋值必须长度和类型必须相同*/


func assignmentArray(t *testing.T)  {
	fmt.Printf("数组之间的相互转化")
	arrayInt := [...]int{1,2,3,4,5}
	newArrayInt := arrayInt
	fmt.Printf("将arrayInt 赋值给newArrayInt: %d\n",newArrayInt)
	arrayInt5 := [5]int{1,2,3,4,5}
	var newArrayInt5 [5]int
	newArrayInt5 = arrayInt5
	fmt.Printf("数组赋值:%d\n",newArrayInt5)
	assert.Equal(t,[...]int{1,2,3,4,5},newArrayInt)
}
/*********************************************************************************************************************/
func transmitArray(t *testing.T)  {
	fmt.Printf("函数之间传递数组")
	arrayInt := [...]int{1,2,3,4,5}
	b := modify(arrayInt)
	fmt.Printf("函数之间传递值之后的值为: %d\n",b)
	assert.Equal(t, [...]int{1,2,3,4,5}, arrayInt)
}

func modify( a [5]int)  (b [5]int){
	a[1] = 3
	return a
}

/*********************************************************************************************************************/


func pointerTransmitArray(t *testing.T)  {
	fmt.Printf("函数之间传递数组")
	arrayInt := [...]int{1,2,3,4,5}
	pointerModify(&arrayInt)
	fmt.Printf("数组开始的指针为:%d\n",arrayInt)
	assert.Equal(t,[...]int{1,3,3,4,5},arrayInt)
}

func pointerModify( a *[5]int){
	a[1] = 3
	fmt.Printf("传递的指针为:%d\n",*a)
}