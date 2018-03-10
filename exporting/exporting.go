package exporting

import (
	"fmt"

	"github.com/john-deng/go-tour/exporting/counters"
	"github.com/john-deng/go-tour/exporting/users"
)

func init()  {
	fmt.Println("main.init()")
}

func changeUser(u users.IUser) {
	u.ChangeName("Mike")
	u.ChangeEmail("mike@outlook.com")
}

func buildUser() users.User {
	user := users.User{"John", "john@email.com"}
	fmt.Println(user)
	return user
}

func Run() {
	fmt.Println(">>> Test exporting ...")

	counter := counters.AlertCounter(10)
	newCounter := counters.New(20)

	fmt.Println(counter)
	fmt.Println(newCounter)

	user := buildUser()

	changeUser(&user)
	fmt.Println(user)

}

