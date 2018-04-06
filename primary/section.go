package primary

import (
	"fmt"
)

/*切片*/

/*切片是一种数据结构，它和数组非常相似，因为他是围绕动态数组的概念设计的，
可以按照需求自动改变大小。使用这种结构，可以更方便的管理和使用数据集合
*/

/*******************************************************************************************************************/
func statementSection()  {
	slice := make([]int,5,5)
	slice =[] int{1,1,1,1,1,1,1,1,1,1}
	fmt.Printf("切片的长度:%v\n",slice)

}


/****************************************************************************************************************/
//nil切片

func nilStatementSection()  {
	//nil切片
	var nilSlice [] int
	//空切片
	slice :=[]int{}
	fmt.Printf("nil切片: %v\n",nilSlice)
	fmt.Printf("空切片： %v\n",slice)
}

/****************************************************************************************************************/
func transmit()  {


	slice := []int{1,2,3,4,5}
	slice1 := slice[:]
	//从0开始索引到最后包含0索引
	slice2 := slice[0:]
	//以0开始5索引结束    不包含5索引
	slice3 := slice[:5]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

}


/****************************************************************************************************************/
func Transmit1()  {


	slice := []int{1,2,3,4,5}
	newSlice := slice[1:3]
	fmt.Printf("%v\n",newSlice)

	fmt.Printf("newSlice长度为: %d\n , newSlice容量:%d\n", len(newSlice) , cap(newSlice))


}
/****************************************************************************************************************/
/*使用切片*/
func Transmit2()  {


	slice := []int{1,2,3,4,5}
	fmt.Printf("%d\n",slice[2])
	slice[2] =10
	fmt.Printf("%d\n",slice[2])

}

/****************************************************************************************************************/
/*当原始数组的容量足够时 不会新建切片原始切片和新切片的值都会发生改变

原始容量不够时才会新建容量
当容量小于1000时总是成倍增长
当容量大于1000时容量增长至原来的25%
*/
func Transmit3()  {
	slice := []int{1,2,3,4,5}
	newslice := slice[1:3]
	fmt.Println("原始切片为",newslice)
	newslice = append(newslice,10)
	fmt.Println("使用后的切片为",newslice)
	fmt.Println("使用后的切片为",slice)
}

/****************************************************************************************************************/

func transmit4()  {
	slice := []int{1,2,3,4,5}
	newslice := slice[1:2:3]
	fmt.Printf("newslice: %d\n",newslice)
	newslice = append(newslice,slice...)
	fmt.Printf("newslice: %d\n",newslice)
}

/****************************************************************************************************************/
func forSlice()  {
	slice := []int{1,2,3,4,5}
	for i,v:= range slice{
		fmt.Printf("索引: %d,值：%d",i,v)
	}
	fmt.Println()
	for _,v:= range slice{
		fmt.Printf("值：%d",v)
	}
	fmt.Println()
	for i :=0; i<len(slice);i++{
		fmt.Printf("值：%d",slice[i])
	}



}





/****************************************************************************************************************/