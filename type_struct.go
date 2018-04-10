package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
func main() {

	/*go语言的方法：方法能给用户定义的类型添加新的行为，方法实际上也是函数，只是在声明时，在关键字func 和方法名之间增加了一个参数*/
	/*关键字func 和函数名之间的参数被称作接收者，将函数与接收者的类型绑在一起。如果一个函数有接收者，这个函数就被称为方法*/

	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("test@email.com")  //内部实现机制 (&bill).changeEmail()

	// 指向user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()   //内部的实现机制 (*lisa).notify()
}
