package main

import "fmt"

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func printSSlice(x []string){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {

	/*申明切片方式：1. 通过make 了申请；2.通过切片字面量来申请；3.使用索引创建切片*/
	/*如果只指定长度， 那么容量和长度相等，都为5*/
	slice := make([]string,5)
	printSSlice(slice)

	/*使用长度和容量申明整型切片， 长度为3， 容量为5， 剩下的2个可以自动扩展*/
	slice1 := make([]int,3,5)
	printSlice(slice1)

	/*2.通过切片字面量来申明，申明切片和申明 数组的区别， 就是数组里面指定了长度*/
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	printSSlice(slice3)


	/*3. 使用索引声明切片*/
	slice4 := []string{99:""}
	printSSlice(slice4)

	/*申请数组*/
/*	slice5 :=[3]string{"Red","Blue","Green"}
	printSSlice(slice5)*/

	/*声明空切片, 空切片的容量和长度都为0*/
	slice6 := make([]int,0)
	printSlice(slice6)

	slice7 :=[]int{}
	printSlice(slice7)


	/*子切片， [lower-bound：upper-bound]获取子切片起， go中 切片共享内存, 修改切片的值会影响原切片的值*/
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)
	testSlice := numbers[1:4]
	testSlice[0] = 5
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println(numbers)

	/*切片的增长,创建一个新的切片*/
	slice8 := []int{10,20,30,40,50}
	newSlice := slice8[1:3]  /*长度为2， 容量为4*/
	newSlice = append(newSlice,60) /*增加一个新元素， 长度为3 < 容量4*/
	fmt.Println(slice8) /*输出[10,20,30,60,50]*/

	newSlice2 := slice8[4:5] /*长度为1， 容量为1*/
	newSlice2 = append(newSlice2,10) /*这里长度 2 > 容量1， 如果切片的底层数组没有足够的可用容量，append 函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新的值*/
	fmt.Println(newSlice2)
	fmt.Println(slice8) /*输出[10,20,30,60,50]*/

	/*使用索引创建切片 slice[i:j:k], 长度: j – i, 容量: k – i*/
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice9 := source[2:3:5]
	fmt.Println(slice9)
	/*end 使用切片创建切片*/

	/*切片for循环迭代*/
	// 创建一个整型切片
	// 其长度和容量都是4 个元素
	slice10 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址，range 创建了每个元素的副本，不是指向元素的所有
	for index, value := range slice10 {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}





}
