package user

import "fmt"

type Users struct {
	Users []User
}

type User struct {
	Name   string
	Age    uint8
	Height float32
}

func MakeUser(users Users) (User, Users) {
	var name string
	var age uint8
	var height float32

	fmt.Println("Введи имя")
	fmt.Scanln(&name)

	fmt.Println("Введи возраст")
	fmt.Scan(&age)

	fmt.Println("Введи рост")
	fmt.Scan(&height)

	var user = User{name, age, height}

	users.Users = append(users.Users, user)

	return user, users
}
