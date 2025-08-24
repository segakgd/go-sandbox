package main

import (
	//"fmt"
	//"go-sanbox/book"
	//"go-sanbox/sega_cmd"
	"go-sanbox/user"
)

func main() {
	user.ConnectAndCreateUser()

	//var userD user.User
	//var users = user.Users{}
	//var books = book.Books{}
	//
	//ok, err := sega_cmd.ChoiceCmd("Хочешь зарегистрироваться?")
	//
	//if err != nil {
	//	fmt.Println("Ошибка:", err)
	//} else if ok {
	//	userD, users = user.MakeUser(users)
	//} else {
	//	userD, users = user.MakeUser(users)
	//}
	//
	//user.ViewUsers(users)
	//
	//fmt.Printf("%s, спасибо, что ты мне написал. Тебе %d лет, а твой рост: %.2f см\n", userD.Name, userD.Age, userD.Height)
	//
	//books = book.AddBook(books)
	//
	//book.ViewBooks(books)
}
