// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)



// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
/*方法集定义了接口的接受规则：
Methods Receivers             values
(t T)                          T *T
(t *T)                          *T
如果使用指针接收者来实现一个接口， 那么只有指向那个类型的指针才能够实现对应的接口；
如果使用值接收者来实现一个接口， 那么 那个类型的值和指向都能够实现对应的接口
*/
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	sendNotification(u)  //编译不过，需要sendNotification(&u)

	// ./go_interface.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
