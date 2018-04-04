package main

import "fmt"

func main() {
	/*创建和初始化map: 1. 使用make:var map_variable map[key_data_type]value_data_type; */
	dict := make(map[string]int) //key 为string ,value 为int
	dict["Red"] =1
	dict1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}  //初始化一个map
	fmt.Println(dict1["Red"])

	/*遍历dict1*/
	for key,value := range dict{
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	/*声明一个未初始化的映射来创建一个值为nil 的映射*/
	var colors map[string]string  /*nil map, 不能赋值*/
	colors["Red"] = "#da1337" /*报错：panic: assignment to entry in nil map*/

}
