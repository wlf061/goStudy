package main

import "fmt"

const MAX int = 3

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)

	return avg
}

func main() {
	/*start 数组的定义，申明, 获取*/
	var balance [6]float32
	balance = [6]float32{10.0, 20.0, 30.0, 12.0, 12.0, 30.0}
	fmt.Println(balance[0])

	var balance1 = []float32{10.0, 20.0, 30.0, 12.0, 12.0, 30.0}
	fmt.Println(balance1[0])

	/*申明数组使用特定索引*/
	array1 := [5]int{1: 10, 2: 20}
	fmt.Println(array1[0])

	/*end 定义，申明，获取数组*/

	/*start 指针数组：数组中存放的几个指针*/
	address := []int{10, 20, 300}
	var ptrAdd [MAX]*int
	var i int
	for i = 0; i < MAX; i++ {
		ptrAdd[i] = &address[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d", i, *ptrAdd[i])
	}
	/*end 指针数组：*/

	/*start 指针数组：用整型指针初始化索引为0 和1 的数组元素*/
	point1 := [5]*int{0: new(int), 1: new(int)}
	*point1[0] = 10
	*point1[1] = 20
	/*end 指针数组*/

	/*start 指向指针的指针*/
	var intValue int
	var addA *int
	var addB **int //指针的指针

	intValue = 200
	addA = &intValue
	addB = &addA

	/*指针取值，使用*， 都输出intValue 的值*/
	fmt.Printf("Value of a = %d\n", intValue)
	fmt.Printf("Value available at *ptr = %d\n", *addA)
	fmt.Printf("Value available at **ptr = %d\n", **addB)

	/*start 在函数间传递数组*/
	/* an int array with 5 elements */
	var  balanceParam = []int {1000, 2, 3, 17, 50}
	var avgValue float32

	/* pass array as an argument */
	avgValue = getAverage(balanceParam, 5 ) ;
	fmt.Println(avgValue)
	/*end 在函数间传递数组*/

	/*start 在函数间传递指针, 相比在函数间传递数组能更节省内存
	  //在函数间传递数组
	  var array [1e6]int
	  func foo(array [1e6]int) {}
	  foo(array)

	  //在函数间传递指针
	  var array [1e6]int
	  func foo(array *[1e6]int) {}
	  foo(&array)

	  //在函数间传递切片
	  slice := make([]int, 1e6)
	  slice = foo(slice)
	 // 函数foo 接收一个整型切片，并返回这个切片，一个切片需要24 字节的内存：指针字段需要8 字节，长度和容量 字段分别需要8 字节
	  func foo(slice []int) []int {}
	  */

}
